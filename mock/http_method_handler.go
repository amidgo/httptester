package mock

import "net/http"

type HttpMethodNotAllowedHandler struct{}

func NewHttpMethodNotAllowedHandler() *HttpMethodNotAllowedHandler {
	return &HttpMethodNotAllowedHandler{}
}

func (h *HttpMethodNotAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
