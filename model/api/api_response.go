package api

type APIResponse struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}
