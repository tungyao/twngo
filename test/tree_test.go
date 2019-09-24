package test

import (
	"../tnwb"
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	r := tnwb.NewRouter()
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello"))
	})
	_ = r.Listening(":80", r)
}
