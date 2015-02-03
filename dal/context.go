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

func (context *Context) Users() ([]*User, error) {
	rows, err := context.tx.Query("SELECT id, firstname, lastname FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		var id int64
		var firstname string
		var lastname string
		if err := rows.Scan(&id, &firstname, &lastname); err != nil {
			return nil, err
		}

		users = append(users, &User{id, firstname, lastname})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}

func (context *Context) CreateUser(user *User) error {
	if user == nil {
		return errors.New("User is nil")
	}

	validationError := user.Validate()
	if validationError != "" {
		return errors.New(validationError)
	}

	rows, err := context.tx.Query("insert into users (firstname, lastname) values ($1, $2);", user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
