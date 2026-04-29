package model

import (
	"reflect"

	"go.gh.ink/json"
)

type ReturnToken struct {
	Token string `json:"token"`
}

type ReturnLinkID struct {
	LinkID string `json:"linkID"`
}

func init() {
	_ = json.PreheatMany([]reflect.Type{
		reflect.TypeOf(ReturnToken{}),
		reflect.TypeOf(ReturnLinkID{}),
	})
}
