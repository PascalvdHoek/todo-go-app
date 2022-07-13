package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-todo-list/errorhandling"
	"go-todo-list/models"
	"net/http"
)

func GetAlbum(c *gin.Context) {
	var (
		album models.Album
		id    string
	)
	id = c.Param("id")

	err = db.NewSelect().Model(&album).Where("id = ? ", id).Limit(1).Scan(ctx)

	if err != nil {
		errorhandling.HandleError(err)
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func GetAlbums(c *gin.Context) {
	albums := new([]models.Album)

	err = db.NewSelect().Model(albums).Scan(ctx)

	if err != nil {
		errorhandling.HandleError(err)
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var album models.Album

	c.BindJSON(&album)
	fmt.Println(album)

	_, err = db.NewInsert().Model(&album).Exec(ctx)
	if err != nil {
		c.IndentedJSON(errorhandling.HandleError(err))
		return
	}
	c.IndentedJSON(http.StatusCreated, album)

}
