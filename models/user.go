package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:i""`

	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	City     int32     `json:"cityId"`
	CityName string    `json:"cityName"`
}
