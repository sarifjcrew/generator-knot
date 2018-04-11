package controllers

import (
	knot "github.com/eaciit/knot/knot.v1"
)

type TestController struct {
	*BaseController
}

func (c *TestController) ExampleHtml(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputTemplate

	return nil
} 

func (c *TestController) ExampleJson(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

