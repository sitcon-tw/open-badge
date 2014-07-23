package app

import (
  // 3rd-party library
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  // Built-in
  "net/http"
  "fmt"
  // Badge
  "github.com/sitcon-tw/open-badge/badge"
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

func OrgizationAPI(res http.ResponseWriter, req *http.Request, r render.Render) {
  r.JSON(200, badge.Orgization{
    Name: "SITCON",
    Url: "http://sitcon.org",
  })
}
