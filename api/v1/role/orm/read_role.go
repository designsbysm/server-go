package orm

import (
	"github.com/designsbysm/server-go/database"
)

func ReadRole(id int) (database.Role, error) {
	var role database.Role
	err := database.DB.First(&role, id).Error

	return role, err
}