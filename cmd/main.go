package main

import (
	"log"

	"github.com/itsshashank/url-shortener/pkg/handler"
)

func main() {
	r := handler.New()
	log.Fatalln(r.Run())
}
