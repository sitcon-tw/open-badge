package main

import (
  "github.com/go-martini/martini"
)

func handleHelloWorld() string {
  return "SITCON Open Badges - Comming Soon...";
}

func main() {
  m := martini.Classic()
  m.Get("/", handleHelloWorld)
  m.Run()
}
