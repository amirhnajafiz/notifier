package request

type Request struct {
	Topic string `json:"topic" xml:"topic" form:"topic"`
	Msg   string `json:"message" xml:"message" form:"message"`
	ID    string `json:"from" xml:"from" form:"from"`
	Date  string `json:"date" xml:"date" form:"date"`
}
