package model

import "gitlab.com/dpcat237/geomicroservices/geolocation"

// Location defines properties of location
type Location struct {
	IPAddress   string  `json:"ip_address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CountryISO  string  `json:"country_iso"`
	CountryName string  `json:"country_name"`
	City        string  `json:"city"`
}

// LocationFromGRPC converts gRPC Location to Location
func LocationFromGRPC(loDt *geolocation.Location) *Location {
	return &Location{
		IPAddress:   loDt.IpAddress,
		Latitude:    loDt.Latitude,
		Longitude:   loDt.Longitude,
		CountryISO:  loDt.Country.Iso,
		CountryName: loDt.Country.Name,
		City:        loDt.City,
	}
}
