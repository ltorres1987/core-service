package user

import (
	"context"
	"delivery-service/internal/utils"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Implementation of the repository in this service.
type userService struct {
	userRepository UserRepository
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewUserService(r UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

// Implementation of 'GetUsers'.
func (s *userService) GetUsers(ctx context.Context) (*[]UserOut, error) {
	return s.userRepository.GetUsers(ctx)
}

// Implementation of 'GetUser'.
func (s *userService) GetUser(ctx context.Context, userID int) (*UserOut, error) {
	return s.userRepository.GetUser(ctx, userID)
}

// Implementation of 'CreateUser'.
func (s *userService) CreateUser(ctx context.Context, signUp *SignUp) (*UserOut, error) {
	// Create a new user struct.
	user := &User{}

	// Set initialized default data for user:
	user.Name = signUp.Name
	user.PasswordHash = utils.GeneratePassword(signUp.Password)
	user.Application = signUp.Application
	user.CreatedUser = signUp.CreatedUser
	user.CreatedAt = time.Now()
	user.Status = "A"

	// Pass to the repository layer.
	result, err := s.userRepository.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	UserOut := &UserOut{
		ID:          int(insertedID),
		Name:        user.Name,
		Application: user.Application,
		CreatedUser: user.CreatedUser,
		CreatedAt:   user.CreatedAt,
		UpdatedUser: user.UpdatedUser,
		UpdatedAt:   user.UpdatedAt,
		Status:      user.Status,
	}
	return UserOut, err
}

// Implementation of 'UpdateUser'.
func (s *userService) UpdateUser(ctx context.Context, userID int, userUpdate *UserUpdate) (*UserOut, error) {

	// Create a new user struct.
	user := &User{}

	// Check if user exists.
	searchedUser, err := s.userRepository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	if searchedUser == nil {
		return nil, fmt.Errorf("There is no user with this ID!")
	}

	// Set value for 'Modified' attribute.
	user.Name = userUpdate.Name
	user.Application = userUpdate.Application
	user.UpdatedUser = userUpdate.UpdatedUser
	user.UpdatedAt = time.Now()

	// Pass to the repository layer.
	err = s.userRepository.UpdateUser(ctx, userID, user)

	if err != nil {
		return nil, err
	}

	return s.userRepository.GetUser(ctx, userID)
}

// Implementation of 'DeleteUser'.
func (s *userService) DeleteUser(ctx context.Context, userID int, userDelete *UserDelete) error {

	// Create a new user struct.
	user := &User{}

	// Check if user exists.
	searchedUser, err := s.userRepository.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	if searchedUser == nil {
		return fmt.Errorf("There is no user with this ID!")
	}

	// Set value for 'Modified' attribute.
	user.UpdatedUser = userDelete.UpdatedUser
	user.UpdatedAt = time.Now()
	user.Status = "I"

	// Pass to the repository layer.
	err = s.userRepository.DeleteUser(ctx, userID, user)

	if err != nil {
		return err
	}

	return nil
}

func (s *userService) UserSignIn(ctx context.Context, signIn *SignIn) (*Tokens, error) {

	// Get user by user name.
	foundedUser, err := s.userRepository.GetUserByName(ctx, signIn.Name)

	if err != nil {
		return nil, err
	}

	if foundedUser == nil {

		return nil, fmt.Errorf("There is no user with this username!")
	}

	// Compare given user password with stored in found user.
	compareUserPassword := utils.ComparePasswords(foundedUser.PasswordHash, signIn.Password)
	if !compareUserPassword {
		// Return, if password is not compare to stored in database.
		return nil, fmt.Errorf("wrong user name or password")
	}

	credentials := make([]string, 3)
	credentials[0] = strconv.Itoa(foundedUser.ID)
	credentials[1] = foundedUser.Name
	credentials[2] = foundedUser.Application

	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(credentials)
	if err != nil {
		// Return status 500 and token generation error.
		return nil, err
	}

	// Create a new Redis connection.
	connRedis, err := utils.RedisConnection()
	if err != nil {
		// Return status 500 and Redis connection error.
		return nil, err
	}

	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	minutes := time.Minute * time.Duration(minutesCount)
	userName := foundedUser.Name

	// Save refresh token to Redis.
	_, err = connRedis.Set(ctx, userName, tokens.Access, minutes).Result()
	if err != nil {
		// Return status 500 and Redis connection error.
		return nil, err
	}

	token := &Tokens{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	}

	return token, nil
}

func (s *userService) UserSignOut(ctx context.Context, userName string) error {

	// Create a new Redis connection.
	connRedis, err := utils.RedisConnection()
	if err != nil {
		// Return status 500 and Redis connection error.
		return err
	}

	// Get token to Redis.
	_, err = connRedis.Del(ctx, userName).Result()
	if err != nil {
		// Return status 500 and Redis connection error.
		return err
	}

	return nil
}
