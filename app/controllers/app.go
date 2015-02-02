package controllers

import (
	"github.com/WiseBird/gooddoglabs_test/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func renderRestError(c *revel.Controller, err error) revel.Result {
	return c.RenderJson(models.RestResult{Error: err.Error()})
}

func renderRestSuccess(c *revel.Controller, data interface{}) revel.Result {
	return c.RenderJson(models.RestResult{Data: data})
}
