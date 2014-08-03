package badge

import (
	"strconv"
	"time"
	"errors"
	"crypto/sha1"
	"crypto/rand"
	"encoding/json"
	"github.com/sitcon-tw/open-badge/badge/storage"
	"github.com/vmihailenco/msgpack"
)

type Badge struct {
	id string
	slug string
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Criteria string `json:"criteria"`
	Issuer string `json:"issuer"`

	// Not support yet.
	Alignment Alignments `json:"alignment,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

func generateBadgeId() string {
	b := make([]byte, 30)
	rand.Read(b)
	return strconv.Itoa(int(time.Now().Unix())) + string(sha1.New().Sum(b))
}

func New(args map[string]string) (*Badge, error) {
	for _, varName := range([]string{"slug", "name", "description", "image", "criteria"}) {
		if v, ok := args[varName]; !ok || v == "" {
			return nil, errors.New("Omit badge " + varName + ".")
		}
	}
	badge := Badge{
		id: generateBadgeId(),
		slug: args["slug"],
		Name: args["name"],
		Description: args["description"],
		Image: args["image"],
		Criteria: args["criteria"],
		Issuer: DefultIssuer.Endpoint(),
	}
	ch := make(chan error)
	data, err := msgpack.Marshal(badge)
    if err != nil {
        return nil, err
    }
	storage.Write("badge", []byte(badge.id), data, ch)
	if err := <- ch ; err != nil {
		return nil, err
	} 
	return &badge, nil
}

func (b *Badge) Update(args map[string]string) (error) {
	if v, ok := args["name"]; ok && v != "" {
		b.Name = args["name"]
	}
	if v, ok := args["description"]; ok && v != "" {
		b.Description = args["description"]
	}
	if v, ok := args["image"]; ok && v != "" {
		b.Image = args["image"]
	}
	if v, ok := args["criteria"]; ok && v != "" {
		b.Criteria = args["criteria"]
	}

	ch := make(chan error)
	data, err := msgpack.Marshal(b)
    if err != nil {
        return err
    }
	storage.Write("badge", []byte(b.id), data, ch)
	return <- ch
}

func (b Badge) EncodeJson(siteUrl string) ([]byte, error) {
	output := b
	output.Issuer = siteUrl + output.Issuer
	return json.Marshal(output)
}