package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	DB = make(map[string]Data)
}

func SaveUrlMapping(index string, Url Data) {
	f, err := os.OpenFile("data.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	jsondata, _ := ioutil.ReadAll(f)
	json.Unmarshal(jsondata, &DB)
	DB[index] = Url
	result, _ := json.Marshal(DB)
	_, err = f.Write(result)
	if err != nil {
		log.Println(err)
	}
}

func RetrieveInitialUrl(shortUrl string) (Data, bool) {
	f, err := os.OpenFile("data.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println("Failed to open file", err)
	}
	defer f.Close()
	jsondata, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(jsondata, &DB)
	if err != nil {
		log.Println(err)
	}
	for _, v := range DB {
		if v.ShortUrl == shortUrl {
			return v, true
		}
	}
	return Data{}, false
}
