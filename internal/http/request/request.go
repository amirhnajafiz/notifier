package request

type Request struct {
	Topic string `json:"topic" xml:"topic" form:"topic"`
	Msg   string `json:"message" xml:"message" form:"message"`
}
