package main

import "github.com/astaxie/beego"


type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func (this *MainController) SayHi()  {
	this.Ctx.WriteString("hi to you")
}

// @router /get/one  [get]
func (this *MainController) GetOne() {
	this.Ctx.WriteString("this is a new one")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Router("/hello", &MainController{}, "get:SayHi")
	beego.Run("localhost:11113")

	//beego.Include(&MainController{})
	//beego.Run("localhost:11113")
}