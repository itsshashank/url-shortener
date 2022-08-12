package db

type Request struct {
	URL string `json:"url"`
}

type Data struct {
	ShortUrl string `json:"shorturl"`
	LongUrl  string `json:"longurl"`
}

var DB map[string]Data
