package incident

import (
	"time"

	"gorm.io/gorm"
)

type InPatient struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	FirstName   string
	LastName    string
	DateOfBirth time.Time
}
