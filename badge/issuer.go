package badge

import (
	"strconv"
)

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
		id: 0,
		slug: "default",
		Name: "SITCON",
		Url: "http://sitcon.org",
		Description: "",
		Image: "http://sitcon.org/logo/sitcon.png",
		Email: "contact@sitcon.org",
		RevocationList: "",
	},
)

func (i Issuer) Endpoint() string {
	if i.slug != "" {
		return "/issuer/" + i.slug
	}
	return "/issuer/" + strconv.Itoa(i.id)
}