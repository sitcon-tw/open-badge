package app

import (
  // 3rd-party library
  "github.com/go-martini/martini"
)

func SetupRouter(m *martini.ClassicMartini) {
  m.Get("/", HomePage)

  // Open Badge API
  m.Get("/api/badge/:id", BadgeAPI)
  m.Get("/api/assertion/:id", AssertionAPI)
  m.Get("/api/organization", OrganizationAPI) // Just response SITCON
}
