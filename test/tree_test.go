package test

import (
	"../tnwb"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestTree(t *testing.T) {
	route := tnwb.NewRouter()
	route.Get("/a", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "/a")
	})
	route.Post("/a/a", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "/a/a")
	})
	if err := route.Listening(":81", route); err != nil {
		log.Println(err)
	}

}
