package controllers

import (
	"advert/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ProductDetailsController(c *gin.Context) {
	urlValues := c.Request.URL.Query()
	fmt.Println(urlValues)

	db := models.Database{ConnectionString: os.Getenv("QUERYSTRING")}
	advertObj := db.GetSpecificAdvert(urlValues["advId"][0])

	c.HTML(http.StatusOK, "product-page.html", gin.H{"advert": advertObj})
}
