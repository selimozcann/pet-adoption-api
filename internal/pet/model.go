package pet

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Type    string `json:"type"`
	Adopted bool   `json:"adopted"`
}
