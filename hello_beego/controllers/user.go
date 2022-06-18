package controllers

import "github.com/astaxie/beego"

type UserController struct {
	// 继承了beego.Controller里的所有方法
	beego.Controller
}

// Get 重写Get方法，首字母大写
func (this *UserController) Get() {
	this.Ctx.WriteString("hello world") // 在当前窗口句柄输出字符串，Ctx是获得当前野蛮的句柄
}

/*func (this *UserController) GetInfo() {
	this.Ctx.WriteString("getinfo data sucess")
}
*/

func (this *UserController) GetInfo() {
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("getinfo data, id =" + id)
}
