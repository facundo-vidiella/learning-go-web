package web

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
type response struct {
	Data interface{} `json:"data"`
}


