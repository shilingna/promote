package controllers

import "github.com/astaxie/beego"

type RController struct {
	beego.Controller
}

func (this *RController) Get() {
	username := this.Ctx.Input.Param(":username")
	this.Ctx.WriteString("getinfo data, username=" + username)
}
