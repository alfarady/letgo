package entity

import (
	"github.com/alfarady/letgo/lib"
)

//StockLimitEntity
type StockLimitEntity struct{
	MinLimit int `json:"min_limit"`
	MaxLimit int `json:"max_limit"`
}

//String
func(p StockLimitEntity)String()string{
	return lib.ObjectToString(p)
}