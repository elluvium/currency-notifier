package models

import "time"

//CurrencyState struct
type CurrencyState struct {
	sourceID      string
	updateDate    time.Time
	organizations Organization
}

//Organization for currency
type Organization struct {
}

//OrganizationType for organization
type OrganizationType struct {
}

//Currency entity
type Currency struct {
}

//Region entity
type Region struct {
}

//City entity
type City struct {
}
