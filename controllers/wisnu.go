package controllers

import (
	knot "github.com/eaciit/knot/knot.v1"
)

type WisnuController struct {
	*BaseController
}

func (c *WisnuController) Getdata(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputTemplate

	return nil
} 

func (c *WisnuController) Get(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputTemplate

	return nil
} 

func (c *WisnuController) Post(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

func (c *WisnuController) Put(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

func (c *WisnuController) Delete(k *knot.WebContext) interface{} {
	k.Config.NoLog = true 
	k.Config.OutputType = knot.OutputJson

	return nil
} 

