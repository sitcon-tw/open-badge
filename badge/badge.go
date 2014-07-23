package badge

type Badge struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Image string `json:"image"`
  Criteria string `json:"criteria"`
  Issuer string `json:"issuer"`
  Alignment []Alignment `json:"alignment,omitempty"`
  Tags []string `json:"tags,omitempty"`
}

type Alignment struct {
  Name string `json:"name"`
  Url string `json:"url"`
  Description string `json:"description,omitempty"`
}
