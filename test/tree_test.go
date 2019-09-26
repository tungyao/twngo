package test

import (
	"../tnwb"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func loggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}
func TestHttpServer(t *testing.T) {
	r := tnwb.NewRouter()
	r.Get("/foo", loggingMiddleware(http.HandlerFunc(foo)))
	_ = r.Listening(":81", r)
}
