package request

type Request struct {
	Msg  string `json:"message" xml:"message" form:"message"`
	ID   string `json:"from" xml:"from" form:"from"`
	Date string `json:"date" xml:"date" form:"date"`
}
