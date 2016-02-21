package controllers

import "github.com/revel/revel"

type List struct {
	*revel.Controller
}

func (h List) Index() revel.Result {
	return h.Render()
}