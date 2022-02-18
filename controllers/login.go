/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-19 12:28:13
 * @LastEditTime: 2022-01-22 17:41:43
 */
package controllers

import (
	"encoding/json"
	"fmt"
	"mybeego/models"
	"mybeego/util"
	"mybeego/validate"

	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

// func (c *LoginController) Get() {
// 	result := &ResultJson{
// 		Code: 1,
// 		Msg:  "success get",
// 	}
// 	c.Data["json"] = result
// 	c.ServeJSON()
// }
func (c *LoginController) Post() {

	//获取数据
	login := &validate.LoginValidate{}
	contentType := c.Ctx.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		//application/json 传输json类型数据接收方式
		body := c.Ctx.Input.RequestBody     // 这是获取到的json二进制数据
		err := json.Unmarshal(body, &login) // json解析位结构体
		if err != nil {
			fmt.Println("json.Unmarshal is err:", err.Error())
		}
		// admin.Username = admin.Username
		// admin.Password = admin.Password
	} else {
		//get post from-data form-urlencoded方式
		login.Username = c.GetString("username")
		login.Password = c.GetString("password")

	}

	result := &ResultLogin{Code: 0, Msg: "登录失败"}

	//验证数据
	err := validate.ValidLogin(login)
	if err != nil {
		result.Msg = err.Error()
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	//数据库查询数据
	data, err := models.GetAdminByUsername(login.Username)

	if err != nil {
		result.Msg = "用户不存在" //err.Error()
	} else if data.Password != util.Password(login.Password) {
		result.Msg = "密码错误"
	} else {
		token, err := util.GenerateToken(data.Id, data.Username, 60*24*24)
		if err != nil {
			result.Msg = err.Error()
		} else {
			result.Code = 1
			result.Msg = "登录成功"
			result.Token = token
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}

/* type AdminPost struct {
	Username string `json:"username" valid:"Required"`
	Password string `json:"password" valid:"Required"`
} */

/* func (c *LoginController) Post() {

	contentType := c.Ctx.Request.Header.Get("Content-Type")
	fmt.Println(contentType)
	fmt.Println(c.Ctx.Input.RequestBody)
	//get post from-data form-urlencoded方式
	username := c.GetString("username")
	password := c.GetString("password")
	//post json 方式
	if contentType == "application/json" {
		var admin AdminPost
		//application/json 传输json类型数据接收方式
		body := c.Ctx.Input.RequestBody     // 这是获取到的json二进制数据
		err := json.Unmarshal(body, &admin) // json解析位结构体
		if err != nil {
			fmt.Println("json.Unmarshal is err:", err.Error())
		}
		username = admin.Username
		password = admin.Password
	}
	result := &ResultLogin{Code: 0, Msg: "登录失败"}

	if username == "" {
		result.Msg = "用户名不能为空"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	if password == "" {
		result.Msg = "密码不能为空"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	//查询数据库
	data, err := models.GetAdminByUsername(username)

	if err != nil {
		result.Msg = "用户不存在" //err.Error()
	} else if data.Password != util.Password(password) {
		result.Msg = "密码错误"
	} else {
		//生成token
		token, err := util.GenerateToken(data.Id, data.Username, 60*24*24)
		if err != nil {
			result.Msg = err.Error()
		} else {
			result.Code = 1
			result.Msg = "登录成功"
			result.Token = token
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
} */

//密码加密方式
// func (c *LoginController) password(value string) (v string) {
// 	o1 := sha1.New()
// 	o1.Write([]byte("Biny_"))
// 	has1 := o1.Sum(nil)
// 	v1 := fmt.Sprintf("%x", has1)

// 	data2 := []byte(value)
// 	has2 := md5.Sum(data2)
// 	v2 := fmt.Sprintf("%x", has2)

// 	encrypt := []byte("_encrypt")
// 	has3 := md5.Sum(encrypt)
// 	v3 := fmt.Sprintf("%x", has3)

// 	o2 := sha1.New()
// 	o2.Write([]byte(value))
// 	has4 := o2.Sum(nil)
// 	v4 := fmt.Sprintf("%x", has4)

// 	o5 := sha1.New()
// 	o5.Write([]byte(v1 + v2 + v3 + v4))
// 	has5 := o5.Sum(nil)
// 	v5 := fmt.Sprintf("%x", has5)

// 	return v5
// }
