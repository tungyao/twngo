package main

import (
	"../tnpool"
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "hello world"+r.URL.Path)
}

func main() {
	t := tnpool.NewTask(func() error {
		http.HandleFunc("/", IndexHandler)
		http.HandleFunc("/a/", IndexHandler)
		http.HandleFunc("/b/", IndexHandler)
		http.HandleFunc("/c/", IndexHandler)
		http.HandleFunc("/d/", IndexHandler)
		http.HandleFunc("/e/", IndexHandler)
		http.HandleFunc("/f/", IndexHandler)
		http.HandleFunc("/g/", IndexHandler)
		http.HandleFunc("/h/", IndexHandler)
		http.HandleFunc("/i/", IndexHandler)
		http.HandleFunc("/j/", IndexHandler)
		http.HandleFunc("/k/", IndexHandler)
		http.HandleFunc("/l/", IndexHandler)
		http.HandleFunc("/m/", IndexHandler)
		http.HandleFunc("/n/", IndexHandler)
		return nil
	})
	p := tnpool.NewPool(20)
	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	_ = http.ListenAndServe(":80", nil)

}
