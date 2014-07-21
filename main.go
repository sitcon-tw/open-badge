package main

import (
  // 3rd-party library
  "github.com/go-martini/martini"
  // Badge
  "github.com/sitcon-tw/open-badge/app"
)

func main() {
  m := martini.Classic()
  app.SetupRouter(m)
  m.Run()
}
