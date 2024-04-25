package entity

import(
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
	"github.com/alfarady/letgo/lib"
)

//GetTrackingNumberResult
type GetTrackingNumberResult struct{
	commonentity.Result
	Response TrackingNumberEntity `json:"response"`
}

//String
func(g GetTrackingNumberResult)String()string{
	return lib.ObjectToString(g)
}