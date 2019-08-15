package test

import (
	"../tnwb"
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	router := tnwb.NewTrie()
	//router.Static("./static")
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

	_ = http.ListenAndServe(":80", router)
}
