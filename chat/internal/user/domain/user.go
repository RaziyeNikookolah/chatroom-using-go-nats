package domain

import (
	"errors"
	"regexp"
	"time"
)

type (
	UserID uint
	Email  string
)

func (e Email) IsValid() bool {
	// Define a regex pattern for validating email addresses
	// This pattern checks for a basic structure of local@domain
	regexPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	regex := regexp.MustCompile(regexPattern)

	// Validate the email using the regex
	return regex.MatchString(string(e))
}

type User struct {
	ID           UserID
	FirstName    string
	LastName     string
	Email        Email
	PasswordHash string
	CreatedAt    time.Time
	DeletedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) Validate() error {
	if !u.Email.IsValid() {
		return errors.New("email is not valid")
	}
	return nil
}
