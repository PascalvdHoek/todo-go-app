package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	models "go-todo-list/models"
	"net/http"
)

func GetUser(c *gin.Context) {
	name := c.Query("name")

	user := new(models.User)
	err = db.NewSelect().Model(user).Where("name=?", name).Scan(ctx)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User with name " + name + " not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	err = db.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		panic(err)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetUserAndCity(c *gin.Context) {
	user := new(models.User)
	err = db.NewSelect().Table("users").ColumnExpr("users.id as id, users.name as Name, City.id as City, City.name as CityName").Join("JOIN City ON users.City = City.id").Scan(ctx, &user.ID, &user.Name, &user.City, &user.CityName)

	fmt.Println("id", user.ID, "name", user.Name, "cityID", user.City, "cityName", user.CityName)
	c.IndentedJSON(http.StatusOK, user)

}
