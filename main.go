/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-19 09:40:28
 * @LastEditTime: 2022-01-21 10:12:41
 */
package main

import (
	"encoding/json"
	"mybeego/controllers"
	_ "mybeego/initial"
	_ "mybeego/routers"
	"mybeego/util"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//两个中间件待分离到单独文件
	//cors 跨域支持
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		//AllowCredentials: true,
	}))
	//jwt CheckToken
	var CheckToken = func(c *context.Context) {
		skip := c.Request.URL.Path
		//fmt.Println(skip)
		if skip != "/login" {
			//定义错误信息返回格式
			result := make(map[string]interface{})
			result["code"] = 0
			//获取token
			token := c.Request.Header.Get("Authorization") //Authorization
			result["token"] = token
			if token == "" {
				//错误返回JSON格式信息
				result["msg"] = "token不能为空"
				res, err := json.Marshal(result)
				if err != nil {
					result["msg"] = "json 格式化错误"
				}
				c.ResponseWriter.Header().Set("Content-Type", "application/json")
				c.ResponseWriter.Write(res)
				return
			}
			token = token[7:] //去除"Bearer "token头
			//验证jwt token
			payload, err := util.ValidateToken(token)
			if err != nil {
				//错误返回JSON格式信息
				result["code"] = -1 //-1为登录超时,需要前端判断跳转到登录页面
				result["msg"] = err.Error()
				res, _ := json.Marshal(result)
				c.ResponseWriter.Header().Set("Content-Type", "application/json")
				c.ResponseWriter.Write(res)
				return
			}
			//fmt.Println(token)
			//fmt.Println(payload)
			//登录成功将uid username传给后面控制器供使用
			c.Input.SetData("uid", payload.UserID)
			c.Input.SetData("username", payload.Username)
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, CheckToken)

	//错误处理
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run("127.0.0.1:8090")

}
