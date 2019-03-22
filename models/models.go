package models

import (
	"time"
)

//CurrencyState struct
type CurrencyState struct {
	SourceID      string `json:"source_id"`
	UpdateDate    time.Time `json:"update_date"`
	Organizations Organization `json:"organizations"`
}

//Organization for currency
type Organization struct {
	OrgType    OrganizationType `json:"org_type"`
	Title      string `json:"title"`
	CityID     string `json:"city_id"`
	Address    string `json:"address"`
	Link       string `json:"link"`
	Currencies []Currency `json:"currencies"`
}

//OrganizationType for organization
type OrganizationType struct {
	TypeID   int `json:"type_id"`
	TypeName string `json:"type_name"`
}

//Currency entity
type Currency struct {
	LongName string `json:"long_name"`
	Name     string `json:"name"`
}

//Region entity
type Region struct {
	RegionID   string `json:"region_id"`
	RegionName string `json:"region_name"`
}

//City entity
type City struct {
	CityID   string `json:"city_id"`
	CityName string `json:"city_name"`
}
