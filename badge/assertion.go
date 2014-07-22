package badge

type Assertion struct {
  Uid string `json:"uid"`
  Recipient Identity `json:"recipient"`
  Badge string `json:"badge"`
  Verify Verification `json:"verify"`
  IssuedOn int64 `json:"issuedOn"`
}

type Verification struct {
  Type string `json:"type"`
  URL string `json:"url"`
}
