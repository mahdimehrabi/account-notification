package dto

type Request struct {
	Method string `json:"method"`
	Data   any    `json:"data"`
	ID     string `json:"id"` // response will send to client base on this id
}
