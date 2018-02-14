package test

import (
	"IPchecker/types"
)

type Test struct {
	Counter int
}

func (item *Test) GetCountryName(ip string) *types.ResolverResponse {
	r := types.ResolverResponse{
		CountryName: "NewerLand",
		IP:          ip,
		HIT:         true,
		CountPerMin: 0,
		CountTotal:  item.CounterInc(),
	}
	return &r
}

func (item *Test) CounterInc() int {
	item.Counter++
	return item.Counter
}
