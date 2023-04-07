package request

type APIRequest struct {
	Request interface{} `json:"request"`
}
type APIResponse struct {
	Response interface{} `json:"response"`
}
type APIResponseHandler struct {
	Responce interface{} `json:"responce"`
}
