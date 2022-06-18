package routers

import (
	"github.com/astaxie/beego"
	"hello_beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// 第三个参数的意思是调用的是GetInfo方法来处理get
	// beego.Router("/user", &controllers.UserController{}, "get:GetInfo")
	beego.Router("/user/?:id", &controllers.UserController{}, "get:GetInfo")
	beego.Router("/api/?:username[\\w]+", &controllers.RController{})
}
