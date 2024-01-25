package main

import (
	"advert/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./views/*")
	router.Static("/static", "./static")

	router.Use(gin.Logger())

	router.GET("/", controllers.IndexController)
	router.GET("/details", controllers.ProductDetailsController)
	router.GET("/addAdvert", controllers.GetCategoriesController)
	router.POST("/addAdvert", controllers.AddAdvPostController)

	router.POST("/api/addAdvert", controllers.ApiAddAdvert)
	router.GET("/api/getAdverts", controllers.ApiGetEntries)

	router.Run("127.0.0.1:3000")
}
