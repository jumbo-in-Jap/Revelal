package controllers

import "github.com/revel/revel"

type User struct {
	*revel.Controller
}

func (h User) Index() revel.Result {
	return h.Render()
}