package main

import (
	"G13.com/backend/configs"
	"G13.com/backend/controller"
	"github.com/gin-gonic/gin"
)

const PORT = "8088"

func main() {
	configs.ConnectDatabase()
	configs.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		// Member routes
		router.GET("/categories", controller.FindCategories)

		// keygame routes
		router.GET("/keygame", controller.FindKeyGame)
		router.POST("/new-keygame", controller.CreateKeyGame)

		// game routes
		router.POST("/new-game", controller.CreateGame)
		router.GET("/game", controller.FindGames)

		//minimumspec routes
		router.POST("/new-minimumspec", controller.CreateMinimumSpec)
		router.GET("/minimumspec", controller.FindMinimumSpec)
	}
	r.Run("localhost: " + PORT)
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
