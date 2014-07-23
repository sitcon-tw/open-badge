package badge

type Assertion struct {
  Uid string `json:"uid"`
  Recipient Identity `json:"recipient"`
  Badge string `json:"badge"`
  Verify Verification `json:"verify"`
  IssuedOn int64 `json:"issuedOn"`
  Image string `json:"image,omitempty"`
  Evidence string `json:"evidence,omitempty"`
  Expires int64 `json:"expires,omitempty"`
}

type Verification struct {
  Type string `json:"type"`
  URL string `json:"url"`
}
