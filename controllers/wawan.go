package controllers

import (
	knot "github.com/eaciit/knot/knot.v1"
)

type WawanController struct {
	*BaseController
}

func (c *WawanController) Index(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

func (c *WawanController) Get(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputTemplate

	return nil
} 

func (c *WawanController) Post(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

func (c *WawanController) Put(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

func (c *WawanController) Delete(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

