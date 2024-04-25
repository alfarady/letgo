package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//CompleteVideoUploadResult
type CompleteVideoUploadResult struct{
	commonentity.Result
	Error string `json:"error"`
	Warning string `json:"warning"`
}

//String
func(g CompleteVideoUploadResult)String()string{
	return lib.ObjectToString(g)
}