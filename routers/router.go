/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-19 09:40:28
 * @LastEditTime: 2022-01-20 20:12:29
 */
package routers

import (
	"mybeego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "post:Post")

	beego.Router("/admin", &controllers.AdminController{}, "get:List")
	beego.Router("/admin/?:id", &controllers.AdminController{}, "get:Get")
	//beego.Router("/admin", &controllers.AdminController{}, "get:List,post:Create,put:Update,delete:Delete")

	beego.Router("/admin", &controllers.AdminController{}, "post:Create")
	//beego.Router("/admin/?:id", &controllers.AdminController{}, "put:Update")
	beego.Router("//admin/?:id", &controllers.AdminController{}, "put:Update")
	beego.Router("/admin", &controllers.AdminController{}, "put:Update")
	beego.Router("/admin", &controllers.AdminController{}, "delete:Delete")
	beego.Router("/auth", &controllers.AdminController{}, "get:Auth")

	// ns := beego.NewNamespace("/v1",
	// 	beego.NSRouter("/admin/?:id", &controllers.AdminController{}, "get:Get"),
	// 	beego.NSRouter("/admin", &controllers.AdminController{}, "get:List"),
	// 	beego.NSRouter("/admin", &controllers.AdminController{}, "post:Create"),
	// 	beego.NSRouter("/admin", &controllers.AdminController{}, "put:Update"),
	// 	beego.NSRouter("/admin", &controllers.AdminController{}, "delete:Delete"),
	// )
	// beego.AddNamespace(ns)
}
