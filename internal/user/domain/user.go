package domain

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"regexp"
	"time"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/conv"
	"github.com/google/uuid"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
)

type (
	UserID   uuid.UUID
	Email    string
	Username string
	Password string
)

type User struct {
	ID        UserID
	CreatedAt time.Time
	DeletedAt time.Time
	Username  Username
	Email     Email
	Password  Password
}

func (u *User) Validate() error {
	if !u.Email.IsValid() {
		return errors.New("Email is not valid")
	}
	return nil
}
func (u UserID) ConvStr() string {
	return uuid.UUID(u).String()
}

func (u *User) PasswordIsCorrect(pass string) bool {
	return NewPassword(pass) == u.Password
}

func NewPassword(pass string) Password {
	h := sha256.New()
	h.Write(conv.ToBytes(pass))
	return Password(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}
func (e Email) IsValid() bool {
	regexPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	regex := regexp.MustCompile(regexPattern)

	// Validate the email using the regex
	return regex.MatchString(string(e))
}
