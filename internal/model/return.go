package model

type ReturnToken struct {
	Token string `json:"token"`
}

type ReturnCNID struct {
	OK bool `json:"ok"`
}

type ReturnTime struct {
	Timestamp   float64 `json:"timestamp"`
	RFC1123     string  `json:"rfc1123"`
	RFC1123Z    string  `json:"rfc1123z"`
	RFC3339     string  `json:"rfc3339"`
	RFC822      string  `json:"rfc822"`
	RFC822Z     string  `json:"rfc822z"`
	RFC850      string  `json:"rfc850"`
	RFC3339Nano string  `json:"rfc3339nano"`
}

type ReturnGreyFilterAccurateSlot struct {
	Description string `json:"description"`
	Object      string `json:"object"`
	Begin       int64  `json:"begin"`
	End         int64  `json:"end"`
}

type ReturnGreyFilterDaySlot struct {
	Description string `json:"description"`
	Object      string `json:"object"`
	Begin       []uint `json:"begin"`
	End         []uint `json:"end"`
}

type ReturnGreyFilter struct {
	AccurateSlot []ReturnGreyFilterAccurateSlot `json:"accurateSlot"`
	DaySlot      []ReturnGreyFilterDaySlot      `json:"daySlot"`
}

type ReturnMusic404 struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Audio  string `json:"audio"`
	Cover  string `json:"cover"`
}

type ReturnMusic404List struct {
	List []ReturnMusic404 `json:"list"`
}

type ReturnFurPassCase struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	BeginDate string `json:"beginDate"`
	EndDate   string `json:"endDate"`
	Location  string `json:"location"`
	Cover     string `json:"cover"`
	Link      string `json:"link"`
}

type ReturnFurPassCaseList struct {
	List []ReturnFurPassCase `json:"list"`
}
