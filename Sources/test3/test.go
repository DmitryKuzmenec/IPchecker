package test3

import (
	"IPchecker/types"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func Init(req chan string, res chan types.ResolverResponse) {
	go func() {
		shaper := time.Tick(100 * time.Second)
		for {
			select {
			case ip := <-req:
				spew.Dump("test3")
				q := types.ResolverResponse{IP: ip, Source: "test3"}
				res <- q
			}
			<-shaper
		}
	}()
}
