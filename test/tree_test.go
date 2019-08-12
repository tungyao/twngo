package test

import (
	"../tnwb"
	"fmt"
	"net/http"
	"testing"
)

func TestTree(t *testing.T) {
	route := tnwb.NewTrie()

	route.Router("get", "/a", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "/a")
	})
	route.Router("get", "/a/a", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "/a/a")
	})
	route.Listening(":80", "", "", route)

}
