package models

import "github.com/uptrace/bun"

type Album struct {
	bun.BaseModel `bun:"table:albums,alias:a""`

	ID     int32   `json: "id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}
