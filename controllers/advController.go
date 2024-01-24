package controllers

import (
	"advert/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCategoriesController(c *gin.Context) {
	db := models.Database{ConnectionString: "user=postgres password=P@ssw0rd! dbname=advboarddb sslmode=disable"}
	categories := db.GetCategories()

	c.HTML(http.StatusOK, "adding-page.html", gin.H{"categories": categories})
}

func AddAdvPostController(c *gin.Context) {
	db := models.Database{ConnectionString: "user=postgres password=P@ssw0rd! dbname=advboarddb sslmode=disable"}

	category := c.PostForm("category")
	title := c.PostForm("title")
	shortDescription := c.PostForm("short-description")
	description := c.PostForm("description")
	imageUrl := c.PostForm("image-url")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 32)
	contact := c.PostForm("contact")

	dateNow := time.Now().Format(time.DateOnly)

	var categoryId string
	_ = categoryId
	categoryMap := db.GetCategories()

	for key, value := range categoryMap {
		categoryId = key
		if category == value {
			categoryId = key
		}
	}

	newAdvert := models.Advertisement{
		Category: category, Title: title, ShortDescription: shortDescription,
		Description: description, Image: imageUrl,
		Price: float32(price), Contact: contact, PubDate: dateNow,
		CategoryId: categoryId,
	}
	fmt.Println(newAdvert)
	db.AddEntries(newAdvert)
}
