package destination

const (
	StatusOk            = "ok"
	StatusForbidden     = "forbidden"
	StatusInternalError = "internal_error"
	StatusBadRequest    = "bad_request"
)

type Request struct {
	Method string `json:"method"`
	Data   any    `json:"data"`
	ID     string `json:"ID"`
}

type Result struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
	ID     string `json:"ID"`
}
