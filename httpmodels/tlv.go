package httpmodels

type TlvRequest struct {
	Data []byte `json:"data"`
}

type TlvResponse struct {
	Type  string `json:"tipo"`
	Len   int    `json:"largo"`
	Value string `json:"valor"`
}
