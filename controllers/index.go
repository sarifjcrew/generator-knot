package controllers

import (
	knot "github.com/eaciit/knot/knot.v1"
)

type IndexController struct {
	*BaseController
}

func (c IndexController) ExampleHtml(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputTemplate

	return nil
} 

