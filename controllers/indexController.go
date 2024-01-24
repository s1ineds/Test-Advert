package controllers

import (
	"advert/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {
	db := models.Database{ConnectionString: os.Getenv("QUERYSTRING")}
	listOfAdverts := db.GetEntries()
	c.HTML(http.StatusOK, "index.html", gin.H{"model": listOfAdverts})
}
