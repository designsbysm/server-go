package jwt

import (
	"testing"

	"github.com/designsbysm/server-go/database"
	"github.com/google/uuid"
)

func TestEncode(t *testing.T) {
	uuid := uuid.New()
	role := database.Role{
		ID:      1,
		IsAdmin: true,
		Name:    "Test",
	}

	token, err := Encode(uuid, role)

	if err != nil {
		t.Error(err)
	} else if token == "" {
		t.Error("empty token")
	}
}
