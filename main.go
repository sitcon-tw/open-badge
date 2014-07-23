package main

import (
  // 3rd-party library
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  // Badge
  "github.com/sitcon-tw/open-badge/app"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())
  app.SetupRouter(m)
  m.Run()
}
