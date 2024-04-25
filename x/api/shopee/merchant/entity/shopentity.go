package entity

import (
	"github.com/alfarady/letgo/lib"
)

//ShopEntity
type ShopEntity struct{
	ShopID int64 `json:"shop_id"`
}


//String
func(g ShopEntity)String()string{
	return lib.ObjectToString(g)
}