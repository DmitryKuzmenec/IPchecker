package test2

import (
	"IPchecker/types"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func Init(req chan string, res chan types.ResolverResponse) {
	go func() {
		shaper := time.Tick(5 * time.Second)
		for {
			select {
			case ip := <-req:
				spew.Dump("test2")
				q := types.ResolverResponse{IP: ip, Source: "test2"}
				res <- q
			}
			<-shaper
		}
	}()
}
