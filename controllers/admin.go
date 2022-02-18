/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-19 13:14:10
 * @LastEditTime: 2022-01-21 09:45:25
 */
package controllers

import (
	"mybeego/models"
	"mybeego/util"
	"strconv"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	result := &ResultJson{Code: 0}
	sid := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(sid, 10, 64)

	data, err := models.GetAdminById(id)
	//fmt.Println(data)
	if err != nil {
		result.Msg = err.Error()
	} else {
		result.Data = data
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *AdminController) List() {
	///fmt.Println(c.Ctx.Input.GetData("uid"))
	//fmt.Println(c.Ctx.Input.GetData("username"))
	limit, _ := c.GetInt("limit", 10)
	page, _ := c.GetInt("page", 1)
	offset := (page - 1) * limit
	result := &ResultJson{Code: 0}
	resultdata := make(map[string]interface{})
	data, err := models.GetAdminByList(limit, offset)
	count, err2 := models.GetAdminCount()
	if err != nil {
		result.Msg = err.Error()
	} else if err2 != nil {
		result.Msg = err2.Error()
	} else {
		resultdata["data"] = data
		resultdata["total"] = count
		resultdata["current"] = page
		resultdata["limit"] = limit
		result.Data = resultdata
	}

	c.Data["json"] = result
	c.ServeJSON()
}
func (c *AdminController) Create() {
	//注意这个引用写法会不会导致高并发下,添加的数据都用同一个指针内存地址,而导致数据共享覆盖问题
	admin := &models.Admin{}
	admin.Username = c.GetString("username")
	admin.Email = c.GetString("username")
	admin.Password = util.Password(c.GetString("password"))
	admin.Mobile = c.GetString("mobile")
	admin.CreateTime = time.Now().Unix()

	result := &ResultJson{Code: 0, Msg: "添加失败"}
	res := models.AdminCreate(admin)
	if res {
		result.Code = 1
		result.Msg = "添加成功"
	}

	c.Data["json"] = result
	c.ServeJSON()
}
func (c *AdminController) Update() {
	admin := models.Admin{}
	admin.Id, _ = c.GetInt64("id")
	admin.Username = c.GetString("username")
	admin.Email = c.GetString("username")
	admin.Password = util.Password(c.GetString("password"))
	admin.Mobile = c.GetString("mobile")
	admin.UpdateTime = time.Now().Unix()
	result := &ResultJson{Code: 1, Msg: "编辑成功"}

	res := models.AdminUpdate(admin)
	if !res {
		result.Code = 0
		result.Msg = "编辑失败"
	}

	c.Data["json"] = result
	c.ServeJSON()
}
func (c *AdminController) Delete() {
	id, _ := c.GetInt64("id", 0)
	result := &ResultJson{Code: 1, Msg: "删除成功"}
	res := models.AdminDelete(id)
	if !res {
		result.Code = 0
		result.Msg = "删除失败"
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *AdminController) Auth() {
	uid := c.Ctx.Input.GetData("uid")
	result := &ResultJson{Code: 1, Msg: "success"}
	resultdata := make(map[string]interface{})
	admin, _ := models.GetAdminById(uid.(int64))
	role, _ := models.GetAdminRoles(uid.(int64))
	menu, _ := models.GetAllMenu()
	//admin.Role = role
	resultdata["admin"] = &admin
	resultdata["role"] = &role
	resultdata["menu"] = &menu

	result.Data = resultdata
	c.Data["json"] = result
	c.ServeJSON()
}
