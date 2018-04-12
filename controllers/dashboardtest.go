package controllers

import (
	knot "github.com/eaciit/knot/knot.v1"
)

type DashboardtestController struct {
	*BaseController
}

func (c *DashboardtestController) Get(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputTemplate

	return nil
} 

func (c *DashboardtestController) Post(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

func (c *DashboardtestController) Put(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

func (c *DashboardtestController) Delete(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

