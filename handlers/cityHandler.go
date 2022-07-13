package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-todo-list/models"
	"net/http"
)

func GetCity(c *gin.Context) {
	cityName := c.Query("name")
	if cityName == "" {

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Please provide a name in the queryparams for this call"})
		return
	}

	var city models.City
	err = db.NewSelect().Model(&city).Where("name =? ", cityName).Scan(ctx)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "City with name " + cityName + " not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, city)

}

func GetCities(c *gin.Context) {
	var cities []models.City
	err = db.NewSelect().Model(cities).Scan(ctx)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, cities)
}
