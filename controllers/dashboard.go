package controllers

import (
	knot "github.com/eaciit/knot/knot.v1"
)

type DashboardController struct {
	*BaseController
}

func (c *DashboardController) ExampleHtml(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputTemplate

	return nil
} 

func (c *DashboardController) ExampleJson(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

