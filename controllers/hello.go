package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Data["Hello"] = "world"
	c.TplName = "hello.tpl"
}

