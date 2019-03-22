package models

import (
	"time"
)

//CurrencyState struct
type CurrencyState struct {
	SourceID      string
	UpdateDate    time.Time
	Organizations Organization
}

//Organization for currency
type Organization struct {
	OrgType    OrganizationType
	Title      string
	CityID     string
	Address    string
	Link       string
	Currencies []Currency
}

//OrganizationType for organization
type OrganizationType struct {
	TypeID   int
	TypeName string
}

//Currency entity
type Currency struct {
	LongName string
	Name     string
}

//Region entity
type Region struct {
	RegionID   string
	RegionName string
}

//City entity
type City struct {
	CityID   string
	CityName string
}
