package badge

import (
	"net/http"
)

func issuerServeMux() *http.ServeMux {
	h := http.NewServeMux()
	h.HandleFunc("/api/issuer", func(w http.ResponseWriter, req *http.Request) {
		
	})
	return h
}
