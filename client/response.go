package client

type Response struct {
	Code    Integer `json:"code"`
	Message string  `json:"msg"`
	Data    any     `json:"data"`
}
