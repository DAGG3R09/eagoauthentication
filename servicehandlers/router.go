package servicehandlers

import (
	"fmt"
	"net/http"
)

type HttpServiceHandler interface {
	Get(*http.Request) string
	Put(*http.Request) string
	Post(*http.Request) string
}

func methodRouter(p HttpServiceHandler, w http.ResponseWriter, r *http.Request) {

	var response string
	if r.Method == "GET" {
		response = p.Get(r)
	} else if r.Method == "PUT" {
		response = p.Put(r)
	} else if r.Method == "POST" {
		response = p.Post(r)
	}
	fmt.Fprintf(w, response)
}
