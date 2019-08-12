package tnwb

import (
	"log"
	"strings"
)
import (
	"net/http"
)

func (mux *Trie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	me, fun := mux.Find(r.URL.Path)
	if fun == nil || r.Method != me {
		w.Header().Set("application", "text/html")
		_, _ = w.Write([]byte("<h1 style=\"font-size=2000px\">404</h1>"))
	}
	if fun != nil {
		fun(w, r)
	}
}
func (mux *Trie) Router(method string, path string, fun http.HandlerFunc) {
	method = strings.ToUpper(method)
	mux.Insert(method, path, fun)
}
func (mux *Trie) Listening(addr string, cert string, key string, handler http.Handler) {
	if ok := http.ListenAndServe(addr, handler); ok != nil {
		log.Print("发生错误")
	}
}
