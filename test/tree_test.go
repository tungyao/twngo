package test

import (
	"../tnwb"
	"fmt"
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	router := tnwb.NewTrie()
	//router.Static("./static")
	router.Group("/user", func(groups *tnwb.Groups) {
		groups.Get("/me", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintln(writer, "HELLO")
		})
	})
	router.Group("/other", func(groups *tnwb.Groups) {
		groups.Get("/me", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintln(writer, "WORLD")
		})
	})
	router.Get("/a", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("application", "text/html")
		_, _ = writer.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <link rel="stylesheet" href="/static/app.css">
</head>
<body>
<h1>123</h1>
</body>
</html>`))

	})

	_ = router.Listening(":80", router)
}
