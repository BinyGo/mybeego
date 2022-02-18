/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-19 09:40:28
 * @LastEditTime: 2022-01-19 10:30:04
 */
package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
