package models

import "github.com/uptrace/bun"

type City struct {
	bun.BaseModel `bun:"table:city,alias:c""`

	ID          int32  `json: "id"`
	Name        string `json: "name"`
	CountryCode string `json: "countryCode"`
	District    string `json: "district"`
	Population  int32  `json: "population"`
}
