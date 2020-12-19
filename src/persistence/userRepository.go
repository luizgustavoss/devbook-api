package persistence

import (
	"database/sql"
	"devbook/src/models"
	"fmt"
)

// UserRepository persists user data
type UserRepository struct {
	db *sql.DB
}

// Create persists a new user
func (repository UserRepository) Create(user models.User) (uint64, error) {

	stmt, error := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer stmt.Close()

	insert, error := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	id, error := insert.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(id), nil;
}

// ListUsers searches for users with name or nick corresponding to description
func (repository UserRepository) ListUsers(description string) ([]models.User, error) {

	description = fmt.Sprintf("%%%s%%", description)

	resultSet, error := repository.db.Query(
		"select u.id, u.name, u.nick, u.email, u.created_at from users u where u.name LIKE ? or u.nick LIKE ?",
		description, description)
	if error != nil {
		return nil, error
	}
	defer resultSet.Close()

	var users []models.User

	for resultSet.Next() {
		var user models.User
		if error = resultSet.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt, ); error != nil {
			return nil, error
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserById gets a specific user by its id
func (repository UserRepository) GetUserById(id uint64) (models.User, error) {

	var user models.User

	resultSet, error := repository.db.Query(
		"select u.id, u.name, u.nick, u.email, u.password, u.created_at from users u where u.id = ?", id)

	if error != nil {
		return user, error
	}
	defer resultSet.Close()

	if resultSet.Next() {
		if error = resultSet.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt, ); error != nil {
			return user, error
		}
	}

	return user, nil
}

// GetUserByEmail gets a specific user by its email
func (repository UserRepository) GetUserByEmail(email string) (models.User, error) {

	var user models.User

	resultSet, error := repository.db.Query(
		"select u.id, u.name, u.nick, u.email, u.password from users u where u.email = ?", email)

	if error != nil {
		return user, error
	}
	defer resultSet.Close()

	if resultSet.Next() {
		if error = resultSet.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, ); error != nil {
			return user, error
		}
	}

	return user, nil
}

// GetUserPasswordById gets a user's password by id
func (repository UserRepository) GetUserPasswordById(id uint64) (string, error) {

	var password = ""

	resultSet, error := repository.db.Query(
		"select u.password from users u where u.id = ?", id)

	if error != nil {
		return password, error
	}
	defer resultSet.Close()

	if resultSet.Next() {
		if error = resultSet.Scan(&password, ); error != nil {
			return password, error
		}
	}

	return password, nil
}

// Update updates user information in database
func (repository UserRepository) Update(id uint64, user models.User) error {

	stmt, error := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ? ")
	if error != nil {
		return error
	}
	defer stmt.Close()

	if _, error = stmt.Exec(user.Name, user.Nick, user.Email, id); error != nil {
		return error
	}

	return nil
}

// Delete deletes a user with the given id
func (repository UserRepository) Delete(id uint64) error {

	stmt, error := repository.db.Prepare(
		"delete from users where id = ? ")
	if error != nil {
		return error
	}
	defer stmt.Close()

	if _, error = stmt.Exec(id); error != nil {
		return error
	}

	return nil
}

// FollowUser register a user following another user
func (repository UserRepository) FollowUser(followedId uint64, followerId uint64) error {

	stmt, error := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)")

	if error != nil {
		return error
	}
	defer stmt.Close()

	insert, error := stmt.Exec(followedId, followerId)
	if error != nil {
		return error
	}

	_, error = insert.LastInsertId()
	if error != nil {
		return error
	}

	return nil
}

// UnfollowUser removes a register of a user following another user
func (repository UserRepository) UnfollowUser(followedId uint64, followerId uint64) error {

	stmt, error := repository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?")

	if error != nil {
		return error
	}
	defer stmt.Close()

	insert, error := stmt.Exec(followedId, followerId)
	if error != nil {
		return error
	}
	_, error = insert.LastInsertId()
	if error != nil {
		return error
	}

	return nil
}

// GetFollowersForUserId searches followers of a user
func (repository UserRepository) GetFollowersForUserId(followedId uint64) ([]models.User, error) {

	resultSet, error := repository.db.Query(
		`select u.id, u.name, u.nick, u.email, u.created_at 
				from users u join followers f on (u.id = f.follower_id) 
				where f.user_id = ?`,
		followedId)
	if error != nil {
		return nil, error
	}
	defer resultSet.Close()

	var users []models.User

	for resultSet.Next() {
		var user models.User
		if error = resultSet.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt, ); error != nil {
			return nil, error
		}
		users = append(users, user)
	}

	return users, nil
}

// GetFollowedUsersForUserId searches users followed by a user
func (repository UserRepository) GetFollowedUsersForUserId(followerId uint64) ([]models.User, error) {

	resultSet, error := repository.db.Query(
		`select u.id, u.name, u.nick, u.email, u.created_at 
				from users u join followers f on (u.id = f.user_id) 
				where f.follower_id = ?`,
		followerId)
	if error != nil {
		return nil, error
	}
	defer resultSet.Close()

	var users []models.User

	for resultSet.Next() {
		var user models.User
		if error = resultSet.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt, ); error != nil {
			return nil, error
		}
		users = append(users, user)
	}

	return users, nil
}

// UpdateUserPassword updates a user's password
func (repository UserRepository) UpdateUserPassword(userId uint64, password string) error {

	stmt, error := repository.db.Prepare(
		"update users set password = ? where id = ? ")
	if error != nil {
		return error
	}
	defer stmt.Close()

	if _, error = stmt.Exec(password, userId); error != nil {
		return error
	}

	return nil
}

// NewUserRepository factory
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}
