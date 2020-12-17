package persistence

import (
	"database/sql"
	"devbook/src/models"
	"fmt"
)

// persists user data
type UserRepository struct {
	db *sql.DB
}

// Create persists a new user
func (repository UserRepository) Create(user models.User) (uint64, error) {

	stmt, error := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
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

// GetUserById gets a specific user bu its id
func (repository UserRepository) GetUserById(id uint64) (models.User, error) {

	var user models.User

	resultSet, error := repository.db.Query(
		"select u.id, u.name, u.nick, u.email, u.created_at from users u where u.id = ?", id)

	if error != nil {
		return user, error
	}
	defer resultSet.Close()

	if resultSet.Next() {
		if error = resultSet.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt, ); error != nil {
			return user, error
		}
	}

	return user, nil
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


// factory
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}
