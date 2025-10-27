package model

type RequestHistory struct {
	Stamp     int64
	AppID     int64
	Interface string
}

type RequestAddLink struct {
	Link     string `json:"link" validate:"required"`
	Validity *int64 `json:"validity"`
}
