package entity

import(
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
	"github.com/alfarady/letgo/lib"
)

//SetNoteResult
type SetNoteResult struct{
	commonentity.Result
}

//String
func(s SetNoteResult)String()string{
	return lib.ObjectToString(s)
}