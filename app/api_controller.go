package app

import (
  // 3rd-party library
  "github.com/go-martini/martini"
  "net/http"
  "fmt"
)

// TODO: get data from database, and reture json data
func BadgeAPI(res http.ResponseWriter, req *http.Request, params martini.Params) {
  id := params["id"]
  res.Write([]byte(fmt.Sprintf("Request Badge ID: %s", id)))
}

// TODO: get data from database, and reture json data
func AssertionAPI(res http.ResponseWriter, req *http.Request, params martini.Params) {
  id := params["id"]
  res.Write([]byte(fmt.Sprintf("Request Assertion ID: %s", id)))
}

// TODO: fill out SITCON detial information
func OrgizationAPI(res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("Comming soon..."))
}
