package entity

import (
	"github.com/alfarady/letgo/lib"
)

//AddressTypeConfigEntity
type AddressTypeConfigEntity struct{
	AddressID int64 `json:"address_id"`
	AddressType []string `json:"address_type"`
}

//String
func(a AddressTypeConfigEntity)String()string{
	return lib.ObjectToString(a)
}
