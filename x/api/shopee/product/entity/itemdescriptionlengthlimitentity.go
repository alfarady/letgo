package entity

import (
	"github.com/alfarady/letgo/lib"
)

//ItemDescriptionLengthLimitEntity
type ItemDescriptionLengthLimitEntity struct{
	MinLimit int `json:"min_limit"`
	MaxLimit int `json:"max_limit"`
}

//String
func(p ItemDescriptionLengthLimitEntity)String()string{
	return lib.ObjectToString(p)
}