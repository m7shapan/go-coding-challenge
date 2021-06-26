package models

type Fact struct {
	Text   string      `json:"text"`
	Number interface{} `json:"number"`
	Found  bool        `json:"found"`
	Type   string      `json:"type"`
}
