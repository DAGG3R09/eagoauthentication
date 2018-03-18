package main

import (
	"eagoauthentication/servicehandlers"
	"log"
	"net/http"
)

func main() {

	p := servicehandlers.PingHandler{}
	a := servicehandlers.AuthenticateHandler{}

	http.Handle("/ping", p)
	http.Handle("/authenticate", a)

	x := http.ListenAndServe(":8080", nil)
	log.Fatal(x)
}
