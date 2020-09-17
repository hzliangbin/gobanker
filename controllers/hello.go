package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}
// @router /:msg [get]
func (o *HelloController) Get() {
	res := o.Ctx.Input.Param(":msg")
	if res != "" {
		o.Data["json"] = res
	} else {
		o.Data["json"] = "hello world"
	}
	o.ServeJSON()
}

