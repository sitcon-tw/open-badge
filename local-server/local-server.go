package main

import (
	"os"
    "log"
    "path"
    "strings"
    "net/http"
)

func main() {

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    	upath := r.URL.Path
		if !strings.HasPrefix(upath, "/") {
			upath = "/" + upath
			r.URL.Path = upath
		}
		upath = path.Clean(upath)
		if upath == "/" {
			upath = "/index.html"
		}
		if _, err := os.Stat("./.tmp" + upath); err == nil {
			http.ServeFile(w, r, "./.tmp" + upath)
		} else if _, err := os.Stat("./app" + upath); err == nil {
			http.ServeFile(w, r, "./app" + upath)
		} else {
			http.NotFound(w, r)
		}
    })
    if err := http.ListenAndServe(":8010", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}