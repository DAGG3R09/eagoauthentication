package servicehandlers

import (
	"net/http"
)

type AuthenticateHandler struct {
}

func (p AuthenticateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p AuthenticateHandler) Get(r *http.Request) string {
	return "AUTHENTICATE GET Called"
}

func (p AuthenticateHandler) Put(r *http.Request) string {
	return "AUTHENTICATE PUT Called"
}

func (p AuthenticateHandler) Post(r *http.Request) string {
	return "AUTHENTICATE POST Called"
}
