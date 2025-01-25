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
type Message struct {
	Sender    string `json:"sender"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}
