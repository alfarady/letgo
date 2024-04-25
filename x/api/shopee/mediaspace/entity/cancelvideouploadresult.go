package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//CancelVideoUploadResult
type CancelVideoUploadResult struct{
	commonentity.Result
	Error string `json:"error"`
	Warning string `json:"warning"`
}

//String
func(g CancelVideoUploadResult)String()string{
	return lib.ObjectToString(g)
}