package controllers

import (
    "github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

type RestError struct {
    Error string `json:"error"`
}

func renderJsonError(c *revel.Controller, err error) revel.Result {
	return c.RenderJson(RestError{err.Error()})
}