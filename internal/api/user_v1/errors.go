package user_v1

const (
	ErrNotValidPassword     = "password must be at least 8 characters long, contain at least one uppercase and one lowercase character"
	ErrPasswordConfirmation = "password confirmation failed"
	ErrNotValidEmail        = "email string is not valid"
	ErrNotValidUsername     = "username must be at least 2 characters long"
)
