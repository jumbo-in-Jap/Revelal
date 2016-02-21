package controllers

import "github.com/revel/revel"

type Movie struct {
	*revel.Controller
}

func (h Movie) Index() revel.Result {
	return h.Render()
}