package model

type User struct {
	Username string
	Email    string
	Type     bool
	Password string
}

type UserProfile struct {
	ID        int
	Username  string
	Email     string
	Type      bool
	CreatedAt string
}
