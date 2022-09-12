package user

import (
	"context"
	"database/sql"
	"time"
)

// User struct to describe User object.
type User struct {
	ID           int       `db:"id"`
	Name         string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	Application  string    `db:"application"`
	CreatedUser  string    `db:"created_user"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedUser  string    `db:"updated_user"`
	UpdatedAt    time.Time `db:"updated_at"`
	Status       string    `db:"status"`
}

// SignUp struct to describe register a new user.
type SignUp struct {
	Name        string `json:"username" validate:"required,email,lte=255"`
	Password    string `json:"password" validate:"required,lte=255"`
	Application string `json:"application" validate:"required,lte=100"`
	CreatedUser string `json:"created_user" validate:"required,lte=100"`
}

// UserUpdate struct to describe update user.
type UserUpdate struct {
	Name        string `json:"username" validate:"required,email,lte=255"`
	Application string `json:"application" validate:"required,lte=100"`
	UpdatedUser string `json:"updated_user" validate:"required,lte=100"`
}

// UserUpdate struct to describe update user.
type UserDelete struct {
	UpdatedUser string `json:"updated_user" validate:"required,lte=100"`
}

type UserOut struct {
	ID          int       `json:"id" `
	Name        string    `json:"username"`
	Application string    `json:"application"`
	CreatedUser string    `json:"created_user"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedUser string    `json:"updated_user"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"`
}

// SignIn struct to describe login user.
type SignIn struct {
	Name     string `json:"username" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

// Tokens struct to describe tokens object.
type Tokens struct {
	Access  string
	Refresh string
}

// Our repository will implement these methods.
type UserRepository interface {
	GetUsers(ctx context.Context) (*[]UserOut, error)
	GetUser(ctx context.Context, userID int) (*UserOut, error)
	CreateUser(ctx context.Context, user *User) (sql.Result, error)
	UpdateUser(ctx context.Context, userID int, user *User) error
	DeleteUser(ctx context.Context, userID int, user *User) error
	GetUserByName(ctx context.Context, userName string) (*User, error)
}

// Our use-case or service will implement these methods.
type UserService interface {
	GetUsers(ctx context.Context) (*[]UserOut, error)
	GetUser(ctx context.Context, userID int) (*UserOut, error)
	CreateUser(ctx context.Context, signUp *SignUp) (*UserOut, error)
	UpdateUser(ctx context.Context, userID int, userUpdate *UserUpdate) (*UserOut, error)
	DeleteUser(ctx context.Context, userID int, userDelete *UserDelete) error
	UserSignIn(ctx context.Context, signIn *SignIn) (*Tokens, error)
	UserSignOut(ctx context.Context, userName string) error
}
