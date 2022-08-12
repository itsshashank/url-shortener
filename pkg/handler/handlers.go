package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsshashank/url-shortener/pkg/db"
)

func New() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("pkg/templates/*")
	r.GET("/", Home)
	r.POST("/short", CreateShortURL)
	r.GET("/:shorturl", RedirectHandler)
	r.GET("/:shorturl/info", GetURL)
	return r
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "URL SHORTENER",
		"body":  "to shortern url \n pass url as payload to /short endpont \n to visit use the shorturl and to view infomation of shorturl use /info endpoint",
	})
}

func CreateShortURL(c *gin.Context) {
	var json db.Request
	if err := c.BindJSON(&json); err != nil || json.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pass url in the json payload"})
		return
	}
	resp, err := http.Get(json.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Broken Url: ", err)})
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Url is giving ", resp.StatusCode)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Url is giving ", resp.StatusCode)})
		return
	}
	myurl, err := GenerateShortUrl(json.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSONP(http.StatusCreated, gin.H{
		"message":  "short url created successfully",
		"ShortUrl": c.Request.Host + "/" + myurl,
	})
}

func RedirectHandler(c *gin.Context) {
	code := c.Param("shorturl")
	url, ok := GetShortUrl(code)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL Not Found"})
		return
	}
	c.Redirect(http.StatusPermanentRedirect, url.LongUrl)
}

func GetURL(c *gin.Context) {
	code := c.Param("shorturl")
	url, ok := GetShortUrl(code)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL Not Found"})
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"ShortUrl": c.Request.Host + "/" + url.ShortUrl,
		"longUrl":  url.LongUrl,
	})
}
