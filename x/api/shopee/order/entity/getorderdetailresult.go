package entity

import(
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
	"github.com/alfarady/letgo/lib"
)

//GetOrderDetailResult
type GetOrderDetailResult struct{
	commonentity.Result
	Response GetOrderDetailResponse `json:"response"`
	Warning []string `json:"warning"`
}
//String
func(g GetOrderDetailResult)String()string{
	return lib.ObjectToString(g)
}
//GetOrderDetailResponse
type GetOrderDetailResponse struct{
	OrderList []OrderEntity `json:"order_list"`
}

//String
func(g GetOrderDetailResponse)String()string{
	return lib.ObjectToString(g)
}