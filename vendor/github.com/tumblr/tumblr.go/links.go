package tumblr

type Links struct {
	Prev Link `json:"prev"`
	Next Link `json:"next"`
}

type Link struct {
	Href string `json:"href"`
	Method string `json:"method"`
	QueryParams map[string]string `json:"query_params"`
}