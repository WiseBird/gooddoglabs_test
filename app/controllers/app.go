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

type RestResult struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func renderRestError(c *revel.Controller, err error) revel.Result {
	return c.RenderJson(RestResult{Error: err.Error()})
}

func renderRestSuccess(c *revel.Controller, data interface{}) revel.Result {
	return c.RenderJson(RestResult{Data: data})
}
