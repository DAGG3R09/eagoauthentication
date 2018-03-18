package servicehandlers

import (
	"net/http"
)

type PingHandler struct {
}

func (p PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p PingHandler) Get(r *http.Request) string {
	return "GET Called"
}

func (p PingHandler) Put(r *http.Request) string {
	return "PUT Called"
}

func (p PingHandler) Post(r *http.Request) string {
	return "POST Called"
}
