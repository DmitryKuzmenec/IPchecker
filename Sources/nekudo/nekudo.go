package nekudo

import (
	"IPchecker/config"
	"IPchecker/db"
	"IPchecker/types"
	"encoding/json"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type resCountry struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type resLocation struct {
	Radius uint    `json:"accuracy_radius"`
	Lat    float64 `json:"latitude"`
	Long   float64 `json:"longitude"`
}

type resJson struct {
	City     bool        `json:"city"`
	Country  resCountry  `json:"country"`
	Location resLocation `json:"location"`
}

func Init(req chan string, res chan types.ResolverResponse) {
	config.Init()
	max_per_min := viper.GetInt("provider.Necudo")

	//skip initialisation if not exists at config
	if max_per_min <= 0 {
		return
	}

	go func() {

		shaper := time.Tick(100000 / time.Duration(max_per_min) * time.Millisecond)
		db := db.Init()
		for {
			select {
			case ip := <-req:
				if q, err := GetCountryName(ip); err == nil {
					db.Save(q)
					res <- q
				} else {
					req <- ip
				}
			}
			<-shaper
		}
	}()
}

func GetCountryName(ip string) (types.ResolverResponse, error) {
	res := types.ResolverResponse{}

	resp, err := http.Get("http://geoip.nekudo.com/api/" + ip)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	data := resJson{}
	if err_d := json.NewDecoder(resp.Body).Decode(&data); err_d != nil {
		return res, err_d
	}

	res = types.ResolverResponse{
		CountryName: data.Country.Name,
		IP:          ip,
		Source:      "necudo",
	}

	return res, nil
}
