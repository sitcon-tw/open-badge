package app

import (
  // 3rd-party library
  "github.com/go-martini/martini"
)

func SetupRouter(m *martini.ClassicMartini) {
  m.Get("/", HomePage)
}
