package model

import (
	"reflect"

	"go.gh.ink/json"
)

type RequestHistory struct {
	Stamp     int64
	AppID     int64
	Interface string
}

type RequestAddLink struct {
	Link     string `json:"link" validate:"required"`
	Validity *int64 `json:"validity"`
}

func init() {
	_ = json.PreheatMany([]reflect.Type{
		reflect.TypeOf(RequestHistory{}),
		reflect.TypeOf(RequestAddLink{}),
	})
}
