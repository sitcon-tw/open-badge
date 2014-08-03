package badge

import (
	"errors"
	"encoding/json"
)

type Badge struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Criteria string `json:"criteria"`
	Issuer string `json:"issuer"`

	// Not support yet.
	Alignment Alignments `json:"alignment,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

func New(args map[string]string) (*Badge, error) {
	for _, varName := range([]string{"name", "description", "image", "criteria"}) {
		if v, ok := args[varName]; !ok || v == "" {
			return nil, errors.New("Omit badge " + varName + ".")
		}
	}

	var badge Badge{
		Name: args["name"],
		Description: args["description"],
		Image: args["image"],
		Criteria: args["criteria"],
		Issuer: DefultIssuer.Endpoint(),
	}
	return badge, nil
}

func (b Badge) EncodeJson(siteUrl string) ([]byte, nil) {
	output := b
	outpupt.Issuer = siteUrl + outpupt.Issuer
	return json.Marshal(output)
}