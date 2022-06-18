package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello_beego/models"
	_ "hello_beego/models"
	_ "hello_beego/routers"
)

func insertUser() {
	o := orm.NewOrm()
	/*user := models.User{}
	user.Name = "lx"
	id, err := o.Insert(&user)
	if err != nil {
		beego.Info("insert error")
		return
	}
	beego.Info("insert success, id = ", id)*/
	res, err := o.Raw(`insert into user values(?,?)`, 13, "zhangsan").Exec()
	if err == nil {
		//返回执行成功条数
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}

}

func queryUser() {
	o := orm.NewOrm()
	user := models.User{Name: "lx"}
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("query is error")
		return
	}
	beego.Info("query success, user=", user)
}

func updateUser() {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	user.Name = "hhhh"
	_, err := o.Update(&user)
	if err != nil {
		beego.Info("update is error")
		return
	}
	beego.Info("update success")
}

func deleteUser() {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	_, err := o.Delete(&user)
	if err != nil {
		beego.Error("delete error")
	}
	beego.Error("delete success")
}

func main() {
	// insertUser()
	// queryUser()
	// updateUser()
	deleteUser()
	beego.Run()
}
