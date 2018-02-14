package src

import (
	"IPchecker/src/test"
	"IPchecker/types"
)

var SourceItems map[string]types.Resolver

func init() {
	SourceItems = make(map[string]types.Resolver)
	SourceItems["test"] = &test.Test{}
}

func GetCountryName(ip string) *types.ResolverResponse {
	source := Dispatcher()
	return source.GetCountryName(ip)
}

func Dispatcher() types.Resolver {
	return SourceItems["test"]
}
