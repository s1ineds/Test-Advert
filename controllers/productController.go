package controllers

import (
	"advert/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductDetailsController(c *gin.Context) {
	urlValues := c.Request.URL.Query()
	fmt.Println(urlValues)

	db := models.Database{ConnectionString: "user=postgres password=P@ssw0rd! dbname=advboarddb sslmode=disable"}
	advertObj := db.GetSpecificAdvert(urlValues["advId"][0])

	c.HTML(http.StatusOK, "product-page.html", gin.H{"advert": advertObj})
}
