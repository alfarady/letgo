package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//AddModelResult
type AddModelResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	Response AddModelResultResponse `json:"response"`
}

//String
func(g AddModelResult)String()string{
	return lib.ObjectToString(g)
}
//AddModelResultResponse
type AddModelResultResponse struct{
	model []ModelEntity `json:"model"`
}

//String
func(g AddModelResultResponse)String()string{
	return lib.ObjectToString(g)
}