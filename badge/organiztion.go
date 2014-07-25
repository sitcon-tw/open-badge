package badge

type Organization struct {
  Name string `json:"name"`
  Url string `json:"url"`
  Description string `json:"description,omitempty"`
  Image string `json:"image,omitempty"`
  Email string `json:"email,omitempty"`
}
