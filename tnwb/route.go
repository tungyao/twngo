package tnwb

import (
	"errors"
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
func (mux *Trie) Get(path string, fun http.HandlerFunc) {
	mux.Insert(http.MethodGet, path, fun)
}
func (mux *Trie) Head(path string, fun http.HandlerFunc) {
	mux.Insert(http.MethodHead, path, fun)
}
func (mux *Trie) Post(path string, fun http.HandlerFunc) {
	mux.Insert(http.MethodPost, path, fun)
}
func (mux *Trie) Put(path string, fun http.HandlerFunc) {
	mux.Insert(http.MethodPut, path, fun)
}
func (mux *Trie) Delete(path string, fun http.HandlerFunc) {
	mux.Insert(http.MethodDelete, path, fun)
}
func (mux *Trie) Listening(parameter ...interface{}) error {
	if len(parameter) != 2 && len(parameter) != 4 {
		return errors.New("参数错误")
	}
	if len(parameter) == 2 {
		if ok := http.ListenAndServe(parameter[0].(string), parameter[1].(http.Handler)); ok != nil {
			return ok
		}
	} else if len(parameter) == 4 {
		if ok := http.ListenAndServeTLS(parameter[0].(string), parameter[1].(string), parameter[2].(string), parameter[3].(http.Handler)); ok != nil {
			return ok
		}
	}
	return nil
}
