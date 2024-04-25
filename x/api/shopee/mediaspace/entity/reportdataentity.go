package entity

import (
	"github.com/alfarady/letgo/lib"
)
//ReportDataEntity
type ReportDataEntity struct{
	UploadCost int `json:"upload_cost"`
}

//String
func(g ReportDataEntity)String()string{
	return lib.ObjectToString(g)
}