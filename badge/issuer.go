package badge

type Issuer struct {
	id int
	slug string
	Name string `json:"name"`
	Url string `json:"url"`
	Description string `json:"description,omitempty"`
	Image string `json:"image,omitempty"`
	Email string `json:"email,omitempty"`
	RevocationList string `json:"revocationList,omitempty"`
}

const (
	DefultIssuer Issuer{ 
		Name: "SITCON",
		Url: "http://sitcon.org",
		Description: "",
		Image: "",
		Email: "contact@sitcon.org",
		RevocationList: "",
	},
	DefultIssuerEndpoint = "/issuer/default",
)