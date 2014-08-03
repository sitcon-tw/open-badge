package server

import (
	"net/http"
)

func badgeServeMux() *http.ServeMux {
	h := http.NewServeMux()
	h.HandleFunc("/api/badge", func(w http.ResponseWriter, req *http.Request) {
		
	})
	return h
}
