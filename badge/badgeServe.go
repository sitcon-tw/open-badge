package badge

import (
	"strings"
	"net/http"
	"github.com/sitcon-tw/open-badge/badge/badge"
)

func badgeServeMux() *http.ServeMux {
	h := http.NewServeMux()
	h.HandleFunc("/api/badge", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimLeft(strings.TrimPrefix(r.URL.Path, "/api/badge"), "/")
		reqMethod := r.Method
		switch reqMethod {
		case "", "GET":	
			if id == "" {
				badgeGetList(w, r)
			} else {
				badgeGet(w, r, id)
			}
			return
		case "POST":
			if id == "" {
				badgeCreate(w, r)
				return
			}
		case "PUT":
			if id != "" {
				badgeUpdate(w, r, id)
				return
			}
		}
		http.NotFound(w, r)
	})
	return h
}

func badgeGetList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetList"))
}

func badgeGet(w http.ResponseWriter, r *http.Request, id string) {
	if b, err := badge.Get(id); err == nil {
		w.Header().Set("Content-Type", "application/json")
		if data, err := b.EncodeJson(ServerHost); err == nil {
			w.Write(data)
			return
		}
	}
	http.NotFound(w, r)
}

func badgeCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if b, err := badge.New(r.PostForm, DefultIssuer.Endpoint()); err == nil {
		w.Header().Set("Content-Type", "application/json")
		if data, err := b.EncodeJson(ServerHost); err == nil {
			w.Write(data)
			return
		}
	}
	http.NotFound(w, r)
}

func badgeUpdate(w http.ResponseWriter, r *http.Request, id string) {
	
}
