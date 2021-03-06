package user

import (
	"net/http"

	"github.com/designsbysm/server-go/database"
	"github.com/gin-gonic/gin"
)

func list(c *gin.Context) {
	user := database.User{}

	list, err := user.List(database.PreloadRole)
	if err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, list)
}
