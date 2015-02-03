package controllers

import (
	"encoding/base64"
	"errors"
	"github.com/WiseBird/gooddoglabs_test/dal"
	"github.com/revel/revel"
	db "github.com/revel/revel/modules/db/app"
	"net/http"
	"strings"
)

type Accounts struct {
	*revel.Controller
	db.Transactional
}

func (c Accounts) CheckAuth() revel.Result {
	res := checkAuth(c.Controller, c.Transactional)
	if res != nil {
		return res
	}

	return renderRestSuccess(c.Controller, nil)
}

func checkAuth(c *revel.Controller, t db.Transactional) revel.Result {
	username, password, ok := basicAuth(c.Request.Request)
	if !ok {
		return renderRestError(c, errors.New("Missing auth info"))
	}

	ok, err := dal.NewContext(t.Txn).CheckAuth(username, password)
	if err != nil {
		return renderRestError(c, err)
	}

	if !ok {
		return renderRestError(c, errors.New("Incorrect username or password"))
	}

	return nil
}

func basicAuth(r *http.Request) (username, password string, ok bool) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return
	}
	return parseBasicAuth(auth)
}
func parseBasicAuth(auth string) (username, password string, ok bool) {
	if !strings.HasPrefix(auth, "Basic ") {
		return
	}
	c, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}
