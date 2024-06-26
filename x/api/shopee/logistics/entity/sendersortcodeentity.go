package entity

import (
	"github.com/alfarady/letgo/lib"
)

//SenderSortCodeEntity
type SenderSortCodeEntity struct{
	FirstSenderSortCode string `json:"first_sender_sort_code"`
	SecondSenderSortCode string `json:"second_sender_sort_code"`
	ThirdSenderSortCode string `json:"third_sender_sort_code"`
}

//String
func(s SenderSortCodeEntity)String()string{
	return lib.ObjectToString(s)
}
