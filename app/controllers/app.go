package controllers

import (
    "github.com/revel/revel"
    "fmt"
    db "github.com/revel/revel/modules/db/app"
)

type App struct {
	*revel.Controller
	db.Transactional
}

func (c App) Index() revel.Result {
    rows, err := c.Txn.Query("SELECT firstname FROM users")
    if err != nil {
        panic(err)
    }
    defer rows.Close()
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            panic(err)
        }
        fmt.Printf("%s\n", name)
    }
    if err := rows.Err(); err != nil {
        panic(err)
    }
    
	return c.Render()
}
