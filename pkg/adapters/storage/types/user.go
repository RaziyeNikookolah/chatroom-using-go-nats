package types

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"not null;uniqueIndex:user_email_idx"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"not null;uniqueIndex:user_email_idx"`
}
