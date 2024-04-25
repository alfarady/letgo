package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//DeleteAddressResult
type DeleteAddressResult struct{
	commonentity.Result
}

//String
func(d DeleteAddressResult)String()string{
	return lib.ObjectToString(d)
}
