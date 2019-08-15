package test

import (
	"../tnwb"
	"fmt"
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	router := tnwb.NewRouter()
	router.Group("/user", func(groups *tnwb.Groups) {
		groups.Get("/me", func(writer http.ResponseWriter, request *http.Request) {
			_, _ = fmt.Fprint(writer, "HELLO")
		})
		groups.Get("/other", func(writer http.ResponseWriter, request *http.Request) {
			_, _ = fmt.Fprint(writer, "other")
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

	_ = http.ListenAndServe(":8000", router)
}
