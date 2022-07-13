package handlers

import (
	"github.com/gin-gonic/gin"
	"go-todo-list/errorhandling"
	"go-todo-list/models"
	"net/http"
)

func GetCountry(c *gin.Context) {
	var (
		country models.Country
		name    string
	)
	name = c.Query("name")
	if name == "" {
		c.IndentedJSON(errorhandling.HandleError(errorhandling.IncorrectQueryParams))
		return
	}

	err = db.NewSelect().Model(&country).Where("name =?", name).Scan(ctx)
	if err != nil {
		c.IndentedJSON(errorhandling.HandleError(err))
		return
	}
	c.IndentedJSON(http.StatusOK, country)

}

func GetCountries(c *gin.Context) {
	countries := new([]models.Country)

	err = db.NewSelect().Model(countries).Scan(ctx)
	if err != nil {
		c.IndentedJSON(errorhandling.HandleError(err))
		return
	}
	c.IndentedJSON(http.StatusOK, countries)
}
