package badge

type Badge struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Image string `json:"image"`
  Criteria string `json:"criteria"`
  Issuer string `json:"issuer"`
}
