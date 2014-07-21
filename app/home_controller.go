package app

import (
  "net/http"
)

func HomePage(res http.ResponseWriter, req *http.Request) {
  // Force change string to []byte
  res.Write([]byte("SITCON Open Badge - Comming Soon ..."))
}
