package types

type ResolverResponse struct {
	CountryName string
	IP          string
	HIT         bool
	CountPerMin int
	CountTotal  int
	Source      string
}

type Resolver interface {
	GetCountryName(string) *ResolverResponse
}
