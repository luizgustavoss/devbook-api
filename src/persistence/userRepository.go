package persistence

import (
	"database/sql"
	"devbook/src/models"
)

// persists user data
type UserRepository struct {
	db *sql.DB
}

// persists a new user
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


// factory
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}
