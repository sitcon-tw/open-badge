package badge

type Verification struct {
	Type string `json:"type"`
	URL string `json:"url"`
}

const (
	VerifyTypeNone = ""
	VerifyTypeHosted = "hosted"
	VerifyTypeSigned = "signed"
)