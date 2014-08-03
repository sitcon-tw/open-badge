package badge

type Alignment struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Description string `json:"description,omitempty"`
}

type Alignments []Alignment