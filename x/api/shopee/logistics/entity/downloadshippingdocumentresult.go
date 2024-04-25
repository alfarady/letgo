package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//DownloadShippingDocumentResult
type DownloadShippingDocumentResult struct{
	commonentity.Result
	File []byte
}

//String
func(d DownloadShippingDocumentResult)String()string{
	return lib.ObjectToString(d)
}