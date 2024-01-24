package controllers

import (
	"advert/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {
	db := models.Database{ConnectionString: "user=postgres password=P@ssw0rd! dbname=advboarddb sslmode=disable"}
	listOfAdverts := db.GetEntries()
	c.HTML(http.StatusOK, "index.html", gin.H{"model": listOfAdverts})
}
