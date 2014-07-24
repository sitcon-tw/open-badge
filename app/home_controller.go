package app

import (
  // 3rd-party library
  "github.com/martini-contrib/render"
  // Built-in
  "net/http"
)

func HomePage(res http.ResponseWriter, req *http.Request, r render.Render) {
  // Force change string to []byte
  res.Write([]byte("SITCON Open Badge - Comming Soon ..."))
}
