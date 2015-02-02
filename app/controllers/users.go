package controllers

import (
    "github.com/revel/revel"
    db "github.com/revel/revel/modules/db/app"
    "github.com/WiseBird/gooddoglabs_test/app/models"
    "errors"
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
    
    rows, err := c.Txn.Query("SELECT id, firstname, lastname FROM users")
    if err != nil {
        return renderJsonError(c.Controller, err)
    }
    defer rows.Close()
    
    users := make([]*models.User, 0)
    
    for rows.Next() {
        var id int64
        var firstname string
        var lastname string
        if err := rows.Scan(&id, &firstname, &lastname); err != nil {
            return renderJsonError(c.Controller, err)
        }
        
        users = append(users, &models.User{ id, firstname, lastname })
    }
    if err := rows.Err(); err != nil {
        return renderJsonError(c.Controller, err)
    }
    
    return c.RenderJson(users)
}

func (c Users) Create(firstname string, lastname string) revel.Result {
    res := checkAuth(c.Controller)
    if res != nil {
        return res
    }
    
    if firstname == "" {
        return renderJsonError(c.Controller, errors.New("Fill first name"))
    }
    if lastname == "" {
        return renderJsonError(c.Controller, errors.New("Fill last name"))
    }
    
    rows, err := c.Txn.Query("insert into users (firstname, lastname) values ($1, $2);", firstname, lastname)
    if err != nil {
        return renderJsonError(c.Controller, err)
    }
    defer rows.Close()
    
    return c.RenderJson("OK")
}