package controllers

import (
	knot "github.com/eaciit/knot/knot.v1"
)

type DashboardController struct {
	*BaseController
}

func (c *DashboardController) Dashboard(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputTemplate

	return nil
} 

func (c *DashboardController) Get(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

func (c *DashboardController) Post(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

func (c *DashboardController) Test(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputTemplate

	return nil
} 

