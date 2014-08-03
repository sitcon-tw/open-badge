package server

import (
	"net/http"
)

func assertionServeMux() *http.ServeMux {
	h := http.NewServeMux()
	h.HandleFunc("/api/assertion", func(w http.ResponseWriter, req *http.Request) {
		
	})
	return h
}
