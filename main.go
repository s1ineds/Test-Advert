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

	router.POST("/api", controllers.AddAdvert)

	router.Run("127.0.0.1:3000")
}
