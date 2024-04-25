package entity

import (
	"github.com/alfarady/letgo/lib"
)

type OrderEntity struct{
	OrderSn string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
	LogisticsStatus string `json:"logistics_status"`
}

//String
func(g OrderEntity)String()string{
	return lib.ObjectToString(g)
}