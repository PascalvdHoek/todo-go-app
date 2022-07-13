package errorhandling

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var IncorrectQueryParams = errors.New("incorrectQueryParams")

func HandleError(err error) (int, map[string]any) {
	if err == sql.ErrNoRows {
		return http.StatusNotFound, gin.H{"message": "Item not found"}
	} else if err == IncorrectQueryParams {
		return http.StatusBadRequest, gin.H{"message": err.Error()}
	} else {
		return http.StatusBadRequest, gin.H{"message": err.Error()}
	}
}
