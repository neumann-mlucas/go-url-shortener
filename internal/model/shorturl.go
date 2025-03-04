package model

type ShortUrl struct {
	Id     int64  `json:"id"`
	Hash   string `json:"hash"`
	Url    string `json:"url"`
	Active bool
}
