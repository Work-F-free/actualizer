package dto

type RequestParams struct {
	Pagination map[string]uint64 `json:"pagination"`
	Query      map[string]string `json:"query"`
}
