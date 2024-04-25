package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//UploadVideoPartResult
type UploadVideoPartResult struct{
	commonentity.Result
	Error string `json:"error"`
	Warning string `json:"warning"`
}

//String
func(g UploadVideoPartResult)String()string{
	return lib.ObjectToString(g)
}