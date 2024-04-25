package entity

import(
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
	"github.com/alfarady/letgo/lib"
)

//UpdateChannelResult
type UpdateChannelResult struct{
	commonentity.Result
	Response UpdateChannelResponse `json:"response"`
}

//String
func(u UpdateChannelResult)String()string{
	return lib.ObjectToString(u)
}
//UpdateChannelResponse
type UpdateChannelResponse struct{
	ShopID int64 `json:"shop_id"`
	Enabled bool `json:"enabled"`
	Preferred bool `json:"preferred"`
	CodEnabled bool `json:"cod_enabled"`
	LogisticsChannelID int64 `json:"logistics_channel_id"`
}