package models

import "github.com/uptrace/bun"

type Country struct {
	bun.BaseModel `bun:"table:country,alias:c""`

	Code           string   `bun:"id,pk" json: "code"`
	Name           string   `json: "name"`
	Continent      string   `json: "continent"`
	Region         string   `json: "region"`
	Surfacearea    float64  `json: "surfaceArea"`
	Indepyear      *int32   `json: "inDepYear"`
	Population     int32    `json: "population"`
	Lifeexpectancy *float64 `json: "lifeExpectancy"`
	GNP            *float64 `json: "gnp"`
	Gnpold         *float64 `json: "gnpOld"`
	Localname      string   `json: "localName"`
	Governmentform string   `json: "governmentForm"`
	Headofstate    *string  `json: "headOfState"`
	Capital        *int32   `json: "capital"`
	Code2          string   `json: "code2"`
}
