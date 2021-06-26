package models

type FactsRequest struct {
	Search  string `query:"s"`
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
}

func NewFactsRequest() *FactsRequest {
	return &FactsRequest{
		Page:    1,
		PerPage: 10,
	}
}

type Filters struct {
	Search string
	Skip   int64
	Limit  int64
}
