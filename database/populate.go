package database

import (
	"github.com/designsbysm/logger/v2"
	"github.com/spf13/viper"
)

func createRecord(data interface{}, query interface{}) {
	err := DB.FirstOrCreate(data, query).Error
	if err != nil {
		logger.Write(
			logger.LevelError,
			err,
		)
	}
}

func populateDB() {
	roleAdmin := Role{
		Name: "admin",
	}
	createRecord(&roleAdmin, roleAdmin)

	userAdmin := User{
		FirstName:    "Scott",
		LastName:     "Matthews",
		Email:        "scott@designsbysm.com",
		PasswordHash: viper.GetString("db.user.password"),
		RoleID:       roleAdmin.ID,
	}
	createRecord(&userAdmin, userAdmin)

	roleUser := Role{
		Name: "user",
	}
	createRecord(&roleUser, roleUser)

	userUser := User{
		FirstName:    "Bob",
		LastName:     "Smith",
		Email:        "bob@designsbysm.com",
		PasswordHash: viper.GetString("db.user.password"),
		RoleID:       roleUser.ID,
	}
	createRecord(&userUser, userUser)
}