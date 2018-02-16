package db

import (
	"IPchecker/types"
	"log"
	"time"

	"github.com/spf13/viper"
	"github.com/tidwall/buntdb"
)

type DbItem struct {
	DB  *buntdb.DB
	Set chan types.ResolverResponse
}

var Db DbItem

func init() {
	//It's also possible to open a database that does not persist to disk by using :memory: as the path of the file.
	//We can use persistent database on disk thee too
	db, err := buntdb.Open(":memory:")
	db.CreateIndex("ip", "*", buntdb.IndexString)

	//BuntDB can have one write transaction opened at a time, but can have many concurrent read transactions.
	set := make(chan types.ResolverResponse, 1000)

	if err != nil {
		log.Fatal(err)
	}
	Db = DbItem{
		DB:  db,
		Set: set,
	}
	go func() {
		//TTL for cache
		ttl := time.Second * time.Duration(viper.GetInt("cache_ttl"))
		for {
			select {
			case data := <-set:
				if data.CountryName != "" && data.IP != "" {
					db.Update(func(tx *buntdb.Tx) error {
						tx.Set(data.IP, data.CountryName, &buntdb.SetOptions{Expires: true, TTL: ttl})
						return nil
					})
				}
			}
		}
	}()
}

func Init() DbItem {
	return Db
}

func (db *DbItem) Save(data types.ResolverResponse) {
	db.Set <- data
}

func (db *DbItem) Get(ip string) (name string) {
	db.DB.View(func(tx *buntdb.Tx) error {
		val, _ := tx.Get(ip)
		name = val
		return nil
	})
	return
}
