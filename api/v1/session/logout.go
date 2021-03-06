package session

import (
	"errors"
	"net/http"

	"github.com/designsbysm/server-go/database"
	"github.com/gin-gonic/gin"
)

func logout(c *gin.Context) {
	data, ok := c.Get("user")
	if !ok {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, errors.New("missing user data"))
		return
	}
	user := data.(database.User)

	session := database.Session{
		UserID: user.ID,
		Token:  "",
	}
	if err := session.Upsert(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
