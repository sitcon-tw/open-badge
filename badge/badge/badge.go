package badge

import (
	"strconv"
	"time"
	"errors"
	"net/url"
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

func New(args url.Values, issuer string) (*Badge, error) {
	for _, varName := range([]string{"slug", "name", "description", "image", "criteria"}) {
		if v := args.Get(varName); v == "" {
			return nil, errors.New("Omit badge " + varName + ".")
		}
	}
	badge := Badge{
		id: generateBadgeId(),
		slug: args.Get("slug"),
		Name: args.Get("name"),
		Description: args.Get("description"),
		Image: args.Get("image"),
		Criteria: args.Get("criteria"),
		Issuer: issuer,
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

func Get(id string) (*Badge, error) {
	chdata := make(chan []byte)
	cherr := make(chan error)
	storage.ReadKey("badge", []byte(id), chdata, cherr)
	var badge Badge
	select {
	case data := <- chdata:
		if err := msgpack.Unmarshal(data, &badge); err != nil {
            return nil, err
        }
	case err := <- cherr:
		return nil, err
	}
	return &badge, nil
}

func (b *Badge) Update(args url.Values) (error) {
	if v := args.Get("name"); v != "" {
		b.Name = args.Get("name")
	}
	if v := args.Get("description"); v != "" {
		b.Description = args.Get("description")
	}
	if v := args.Get("image"); v != "" {
		b.Image = args.Get("image")
	}
	if v := args.Get("criteria"); v != "" {
		b.Criteria = args.Get("criteria")
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