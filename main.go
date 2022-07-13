package main

import (
	"github.com/gin-gonic/gin"
	"go-todo-list/handlers"
)

func main() {
	router := gin.Default()
	router.Routes()
	router.GET("/user", handlers.GetUser)
	router.GET("/userCity", handlers.GetUserAndCity)
	router.GET("/users", handlers.GetUsers)
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbum)
	router.POST("/albums", handlers.PostAlbum)
	router.GET("/city", handlers.GetCity)
	router.GET("/cities", handlers.GetCities)
	router.GET("/country", handlers.GetCountry)
	router.GET("/countries", handlers.GetCountries)
	router.Run("localhost:8080")
}
