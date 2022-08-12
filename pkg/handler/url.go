package handler

import (
	"encoding/base64"
	"log"
	"net/url"

	"github.com/google/uuid"
	"github.com/itsshashank/url-shortener/pkg/db"
)

func GenerateShortUrl(longurl string) (string, error) {
	u, err := url.ParseRequestURI(longurl)
	if err != nil {
		return "", err
	}
	code := base64.StdEncoding.EncodeToString([]byte(u.Host + u.RequestURI()))
	log.Println(u.Host, u.RequestURI())
	index := code[:4] + code[len(code)-4:]
	item, ok := db.DB[index]
	if ok {
		if ur, _ := url.Parse(item.LongUrl); ur.Host == u.Host {
			return item.ShortUrl, nil
		}
	}
	shorturl := uuid.New().String()[:4] + base64.StdEncoding.EncodeToString([]byte(u.Host))[:3]
	db.SaveUrlMapping(index, db.Data{ShortUrl: shorturl, LongUrl: longurl})
	return shorturl, nil
}

func GetShortUrl(code string) (db.Data, bool) {
	return db.RetrieveInitialUrl(code)
}
