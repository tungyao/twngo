package tnwb

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var FileType = map[string]string{"css": "text/css"}

func writeStaticFile(path string, filename []string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", FileType[filename[1]])
	if f, err := os.Open("." + path); err == nil {
		data, _ := ioutil.ReadAll(f)
		_, _ = w.Write(data)
	}
}
func (mux *Trie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(`^/static/\w+\.\w+$`)
	file := reg.FindStringSubmatch(r.URL.String())
	log.Println(file)
	if len(file) != 0 {
		filename := strings.Split(file[0], ".")
		writeStaticFile(r.URL.Path, filename, w)
		return
	}

	me, fun := mux.Find(r.URL.Path)
	if fun == nil || r.Method != me {
		w.Header().Set("Content-type", "text/html")

		_, _ = w.Write([]byte("<h1 style=\"font-size=2000px;\">404</h1>"))
	}
	if fun != nil {
		fun(w, r)
	}
}
func (mux *Trie) Get(path string, fun http.HandlerFunc) {
	mux.Insert(http.MethodGet, path, fun)
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
func (mux *Trie) Static(filepath string) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}
func (mux *Trie) Listening(parameter ...interface{}) error {
	if len(parameter) != 2 && len(parameter) != 4 {
		return errors.New("parameter length must is 2 or 4")
	}
	if len(parameter) == 2 {
		if ok := http.ListenAndServe(parameter[0].(string), parameter[1].(http.Handler)); ok != nil {
			return ok
		}
	} else {
		if ok := http.ListenAndServeTLS(parameter[0].(string), parameter[1].(string), parameter[2].(string), parameter[1].(http.Handler)); ok != nil {
			return ok
		}
	}
	log.Println("service is running")
	return nil
}
