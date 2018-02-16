package freegeoip

import (
	"IPchecker/config"
	"IPchecker/db"
	"IPchecker/types"
	"encoding/json"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type resJson struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   uint    `json:"metro_code"`
}

func Init(req chan string, res chan types.ResolverResponse) {
	config.Init()
	providers := viper.GetStringMapString("providers")

	//skip initialisation if not exists at config
	if _, ok := providers["freegeoip"]; !ok {
		return
	}

	go func() {
		shaper := time.Tick(1 * time.Second)
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

	resp, err := http.Get("http://freegeoip.net/json/" + ip)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	data := resJson{}
	if err_d := json.NewDecoder(resp.Body).Decode(&data); err_d != nil {
		return res, err_d
	}

	res = types.ResolverResponse{
		CountryName: data.CountryName,
		IP:          ip,
		Source:      "freegeoip",
	}
	return res, nil
}
