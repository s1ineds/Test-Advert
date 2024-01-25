/*

	Invoke-RestMethod -Method Post -Uri "http://localhost:3000/api/addAdvert" -Headers @{"Content-Type"="application/json"} -Body (@{"Id"=1;"Category"="Услуги";"Title"="TestTitle1";"ShortDescription"="TestShortDescription1";"Description"="TestDescription";"Image"="TestImageUrl";"Price"=100;"Contact"="TestContact";"PubDate"="TestPubDate";"CategoryId"="1"}|ConvertTo-Json)

*/

package controllers

import (
	"advert/models"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiAddAdvert(c *gin.Context) {
	var newAdvert models.Advertisement
	db := models.Database{ConnectionString: os.Getenv("QUERYSTRING")}
	categoriesMap := db.GetCategories()

	err := c.BindJSON(&newAdvert)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range categoriesMap {
		if value == newAdvert.Category {
			newAdvert.CategoryId = key
			break
		}
	}

	fmt.Println(newAdvert)
	db.AddEntries(newAdvert)
}
