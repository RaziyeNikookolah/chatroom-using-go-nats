package domain

type User struct {
	Username string
	Password string
	Email    string
}
type (
	Token    string
	password string
	emain    string
)
type UserClaim struct {
	Username string
	ID       string
	Email    string
}
