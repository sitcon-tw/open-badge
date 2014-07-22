package badge

type Identity struct {
  Identity string `json:"identity"`
  Type string `json:"type"`
  Hashed bool `json:"hashed"`
  Salt string `json:"salt"`
}
