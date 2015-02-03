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
	res := checkAuth(c.Controller)
	if res != nil {
		return res
	}

	context := dal.NewContext(c.Txn)

	users, err := context.Users()
	if err != nil {
		return renderRestError(c.Controller, err)
	}

	return renderRestSuccess(c.Controller, users)
}

func (c Users) Create(firstname string, lastname string) revel.Result {
	res := checkAuth(c.Controller)
	if res != nil {
		return res
	}

	user := &dal.User{FirstName: firstname, LastName: lastname}

	context := dal.NewContext(c.Txn)
	err := context.CreateUser(user)
	if err != nil {
		return renderRestError(c.Controller, err)
	}

	return renderRestSuccess(c.Controller, nil)
}
