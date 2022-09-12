package user

import (
	"context"
	"database/sql"
)

// Queries that we will use.
const (
	QUERY_GET_USERS        = "SELECT id, name, application, created_user, created_at, updated_user, updated_at, status FROM users WHERE status = ?"
	QUERY_GET_USER         = "SELECT id, name, application, created_user, created_at, updated_user, updated_at, status FROM users WHERE  id = ? and status = ?"
	QUERY_CREATE_USER      = "INSERT INTO users (name, password_hash, application, created_user, created_at, updated_user, updated_at, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	QUERY_UPDATE_USER      = "UPDATE users SET name = ?, application = ?, updated_user = ?, updated_at = ? WHERE id = ?"
	QUERY_DELETE_USER      = "UPDATE users SET status = ?, updated_user = ?, updated_at = ? WHERE id = ?"
	QUERY_GET_USER_BY_NAME = "SELECT id, name, password_hash, application, created_user, created_at, updated_user, updated_at, status FROM users WHERE  name = ? and status = ?"
)

// Represents that we will use MariaDB in order to implement the methods.
type mariaDBRepository struct {
	mariadb *sql.DB
}

// Create a new repository with MariaDB as the driver.
func NewUserRepository(mariaDBConnection *sql.DB) UserRepository {
	return &mariaDBRepository{
		mariadb: mariaDBConnection,
	}
}

// Gets all users in the database.
func (r *mariaDBRepository) GetUsers(ctx context.Context) (*[]UserOut, error) {
	// Initialize variables.
	var users []UserOut

	// Get all users.
	res, err := r.mariadb.QueryContext(ctx, QUERY_GET_USERS, "A")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	// Scan all of the results to the 'users' array.
	// If it's empty, return null.
	for res.Next() {
		user := &UserOut{}
		err = res.Scan(&user.ID, &user.Name, &user.Application, &user.CreatedUser, &user.CreatedAt, &user.UpdatedUser, &user.UpdatedAt, &user.Status)
		if err != nil && err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}

	// Return all of our users.
	return &users, nil
}

// Gets a single user in the database.
func (r *mariaDBRepository) GetUser(ctx context.Context, userID int) (*UserOut, error) {
	// Initialize variable.
	user := &UserOut{}

	// Prepare SQL to get one user.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_GET_USER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Get one user and insert it to the 'user' struct.
	// If it's empty, return null.
	err = stmt.QueryRowContext(ctx, userID, "A").Scan(&user.ID, &user.Name, &user.Application, &user.CreatedUser, &user.CreatedAt, &user.UpdatedUser, &user.UpdatedAt, &user.Status)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Return result.
	return user, nil
}

// Creates a single user in the database.
func (r *mariaDBRepository) CreateUser(ctx context.Context, user *User) (sql.Result, error) {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_CREATE_USER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Insert one user.
	result, err := stmt.ExecContext(ctx, user.Name, user.PasswordHash, user.Application, user.CreatedUser, user.CreatedAt, user.UpdatedUser, user.UpdatedAt, user.Status)
	if err != nil {
		return nil, err
	}

	// Return empty.
	return result, nil
}

// Updates a single user in the database.
func (r *mariaDBRepository) UpdateUser(ctx context.Context, userID int, user *User) error {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_UPDATE_USER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Update one user.
	_, err = stmt.ExecContext(ctx, user.Name, user.Application, user.UpdatedUser, user.UpdatedAt, userID)
	if err != nil {
		return err
	}

	// Return empty.
	return nil
}

// Deletes a single user in the database.
func (r *mariaDBRepository) DeleteUser(ctx context.Context, userID int, user *User) error {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_DELETE_USER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Delete one user.
	_, err = stmt.ExecContext(ctx, user.Status, user.UpdatedUser, user.UpdatedAt, userID)
	if err != nil {
		return err
	}

	// Return empty.
	return nil
}

// Gets a single user in the database.
func (r *mariaDBRepository) GetUserByName(ctx context.Context, userName string) (*User, error) {
	// Initialize variable.
	user := &User{}

	// Prepare SQL to get one user.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_GET_USER_BY_NAME)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Get one user and insert it to the 'user' struct.
	// If it's empty, return null.
	err = stmt.QueryRowContext(ctx, userName, "A").Scan(&user.ID, &user.Name, &user.PasswordHash, &user.Application, &user.CreatedUser, &user.CreatedAt, &user.UpdatedUser, &user.UpdatedAt, &user.Status)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Return result.
	return user, nil
}
