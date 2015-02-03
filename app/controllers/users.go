package controllers

import (
	"github.com/WiseBird/gooddoglabs_test/dal"
	"github.com/revel/revel"
	db "github.com/revel/revel/modules/db/app"
)

type Users struct {
	*revel.Controller
	db.Transactional
}

func (c Users) List() revel.Result {
	res := checkAuth(c.Controller, c.Transactional)
	if res != nil {
		return res
	}

	users, err := dal.NewContext(c.Txn).Users()
	if err != nil {
		return renderRestError(c.Controller, err)
	}

	return renderRestSuccess(c.Controller, users)
}

func (c Users) Create(username, firstname, lastname, password string) revel.Result {
	res := checkAuth(c.Controller, c.Transactional)
	if res != nil {
		return res
	}

	user := &dal.User{UserName: username, FirstName: firstname, LastName: lastname}

	err := dal.NewContext(c.Txn).CreateUser(user, password)
	if err != nil {
		return renderRestError(c.Controller, err)
	}

	return renderRestSuccess(c.Controller, nil)
}
