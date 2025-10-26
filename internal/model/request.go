package model

type RequestHistory struct {
	Stamp     int64
	AppID     int64
	Interface string
}

type RequestCNID struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type RequestSMSLogtoBasicPayload struct {
	Code string `json:"code" validate:"required"`
}

type RequestSMSLogtoBasic struct {
	To      string                      `json:"to" validate:"required"`
	Type    string                      `json:"type" validate:"required"`
	Payload RequestSMSLogtoBasicPayload `json:"payload" validate:"required"`
}
