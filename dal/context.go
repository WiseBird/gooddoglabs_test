package dal

import (
	"database/sql"
	"errors"
)

type Context struct {
	tx *sql.Tx
}

func NewContext(tx *sql.Tx) *Context {
	return &Context{tx}
}

func (context *Context) CheckAuth(username, password string) (bool, error) {
	rows, err := context.tx.Query("SELECT id FROM users WHERE username = $1 AND password = $2", username, password)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), rows.Err()
}

func (context *Context) Users() ([]*User, error) {
	rows, err := context.tx.Query("SELECT id, username, firstname, lastname FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		var id int64
		var username string
		var firstname string
		var lastname string
		if err := rows.Scan(&id, &username, &firstname, &lastname); err != nil {
			return nil, err
		}

		users = append(users, &User{id, username, firstname, lastname})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}

func (context *Context) CreateUser(user *User, password string) error {
	if user == nil {
		return errors.New("User is nil")
	}

	if password == "" {
		return errors.New("Password is empty")
	}

	validationError := user.Validate()
	if validationError != "" {
		return errors.New(validationError)
	}

	rows, err := context.tx.Query("insert into users (username, firstname, lastname, password) values ($1, $2, $3, $4);",
		user.UserName, user.FirstName, user.LastName, password)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
