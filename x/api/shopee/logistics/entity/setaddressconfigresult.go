package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//AddressTypeConfigEntity
type SetAddressConfigResult struct{
	commonentity.Result
}

//String
func(a SetAddressConfigResult)String()string{
	return lib.ObjectToString(a)
}
