package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mayowa/bjf"
	"github.com/triangletodd/gort/internal/models"
)

type UrlController struct{}

var urlModel = new(models.URL)

func (u UrlController) Handler(c *gin.Context) {
	pathList := strings.Split(c.Request.URL.Path, "/")
	method := c.Request.Method

	if method == "GET" {
		if len(pathList) == 2 && pathList[1] != "" {
			UrlController{}.Redirect(c)
			return
		}
	}

	c.JSON(400, gin.H{"message": "bad request"})
}

func (u UrlController) Redirect(c *gin.Context) {
	short := strings.Split(c.Request.URL.Path, "/")[1]
	urlName := strconv.Itoa(bjf.Decode(short))

	url, err := urlModel.GetByName(urlName)
	if err != nil {
		c.JSON(404, gin.H{"message": "URL not found.", "error": err})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Long)
}

func (u UrlController) RetrieveUrls(c *gin.Context) {
	urls, err := urlModel.List()
	if err != nil {
		c.JSON(500, gin.H{"message": "Could not retrieve URLs.", "error": err})
	}

	c.JSON(200, gin.H{"count": len(urls), "items": urls})
}

func (u UrlController) RetrieveUrl(c *gin.Context) {
	short := c.Param("short")
	urlName := strconv.Itoa(bjf.Decode(short))

	url, err := urlModel.GetByName(urlName)
	if err != nil {
		c.JSON(404, gin.H{"message": "URL not found.", "error": err})
		return
	}

	c.JSON(200, gin.H{"name": url.Name, "url": url.Long})
}

func (u UrlController) CreateUrl(c *gin.Context) {
	urlString := c.PostForm("url")
	url, err := urlModel.Create(urlString)
	if err != nil {
		c.JSON(500, gin.H{"message": "Could not create short URL.", "error": err})
		return
	}

	c.JSON(200, gin.H{"name": url.Name, "url": url.Long, "short": url.Short})
}
