/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-20 16:20:28
 * @LastEditTime: 2022-01-21 09:26:24
 */
package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

type AdminMenu struct {
	Id         int64        `json:"id"`
	ParentId   int64        `json:"parent_id"`
	Type       int          `json:"type"`
	Status     int          `json:"status"`
	Sort       int64        `json:"sort"`
	Controller string       `json:"controller"`
	Action     string       `json:"action"`
	Param      string       `json:"param"`
	Path       string       `json:"path"`
	Title      string       `json:"title"`
	Icon       string       `json:"icon"`
	IsMenu     int          `json:"is_menu"`
	Level      int          `json:"level"`
	Children   []*AdminMenu `json:"children" orm:"-"`
}

func init() {

	orm.RegisterModelWithPrefix("biny_", new(AdminMenu))

}

func GetAllMenu() ([]*AdminMenu, error) {
	o := orm.NewOrm()

	var menus []*AdminMenu
	menu := new(AdminMenu)
	_, err := o.QueryTable(menu).Filter("is_menu", 1).All(&menus) //All(&admins, "Id", "Username")
	menus = makeTree(menus, 0, 0)
	return menus, err
}

// * // 递归实现无限分类
func makeTree(menus []*AdminMenu, pid int64, level int) []*AdminMenu {
	fmt.Println(&menus)
	var tree []*AdminMenu
	for i := 0; i < len(menus); {
		row := menus[i]
		if row.ParentId == pid {
			row.Level = level
			menus = append(menus[:i], menus[i+1:]...)
			fmt.Println(menus)
			children := makeTree(menus, row.Id, level+1)
			if children != nil {
				row.Children = children
			}
			tree = append(tree, row)
		} else {
			i++
		}
	}
	return tree
}
