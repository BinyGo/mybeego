/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-19 10:34:37
 * @LastEditTime: 2022-01-19 19:49:21
 */
package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	beego.Controller
}

type ResultJson struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Token string      `json:"token"`
}

type ResultLogin struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

func (c *ErrorController) Error404() {

	//type ErrorJson map[string]interface{}
	//result := make(map[string]interface{})
	// result := make(ErrorJson)
	// result["code"] = 0
	// result["msg"] = "page not found"
	result := &ResultJson{
		Code: 1,
		Msg:  "page not found",
		//Data: make([]interface{}, 0),
	}
	// result.Code = 1
	// result.Msg = "not"

	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ErrorController) Error() {
	result := &ResultJson{
		Code: 1,
		Msg:  "Method Not Allowed",
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ErrorController) Error501() {
	result := &ResultJson{
		Code: 1,
		Msg:  "server error",
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ErrorController) ErrorDb() {
	result := &ResultJson{
		Code: 1,
		Msg:  "database is now down",
	}
	c.Data["json"] = result
	c.ServeJSON()
}
