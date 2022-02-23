package request

type Request struct {
	Msg string `json:"message" xml:"message" form:"message"`
}
