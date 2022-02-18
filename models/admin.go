/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-19 15:30:01
 * @LastEditTime: 2022-01-23 13:24:09
 */
package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"` //`json:"-" orm:"-" orm:"column(id)`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	Group      int    `json:"group"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type AdminCount struct {
	Id    int64 `json:"id"`
	Count int64 `json:"count"`
}

type AdminRole struct {
	Id      int64 `json:"id"`
	AdminId int64 `json:"admin_id"`
	RoleId  int64 `json:"role_id"`
	//Role *Role `orm:"rel(fk)"`
	//Admin   *Admin `orm:"reverse(one)"`
	//Admin   *Admin `orm:"rel(fk)"`
	//Roles []*Role `orm:"rel(m2m)"` //`json:"roles"` // orm:"reverse(m2m)"
}

type Role struct {
	Id    int64  `json:"id"`
	Group int    `json:"group"`
	Title string `json:"title"`
	Auth  string `json:"auth"`
	//AdminRoles []*AdminRole //`orm:"reverse(many)"`
}

func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	// orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.31.22:3306)/tp6-admin?charset=utf8")
	// fmt.Println("RegisterDataBase")
	orm.RegisterModelWithPrefix("biny_", new(Admin), new(AdminCount), new(AdminRole), new(Role))
	//orm.RegisterModel(new(Admin))

}

func GetAdminByUsername(username string) (Admin, error) {
	db := orm.NewOrm()
	admin := Admin{Username: username}
	err := db.Read(&admin, "Username")
	return admin, err
}

func GetAdminById(id int64) (Admin, error) {
	db := orm.NewOrm()
	admin := Admin{Id: id}
	err := db.Read(&admin)
	return admin, err
}

func GetAdminByList(limit int, offset int) ([]*Admin, error) {
	o := orm.NewOrm()
	//qs = o.QueryTable(&Admin)
	var admins []*Admin
	admin := new(Admin)                                                            //admin := &Admin{} 忘记怎么写的new了,待测试使用这个代码是否可行,应该是官方示例复制的代码,不能变
	_, err := o.QueryTable(admin).OrderBy("-id").Limit(limit, offset).All(&admins) //All(&admins, "Id", "Username")
	//fmt.Println(qs)
	return admins, err
}

func GetAdminCount() (int64, error) {
	db := orm.NewOrm()
	aminCount := AdminCount{Id: 1}
	err := db.Read(&aminCount)
	return aminCount.Count, err

	// aminCount := AdminCount{Id: 1}
	// err := db.Read(&aminCount)
	//return aminCount, err
	// count, err := db.QueryTable("biny_admin_count").Count()
	// return count, err
}

func AdminCreate(data *Admin) bool {
	db := orm.NewOrm()
	//闭包事务
	err := db.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// data
		id, err := txOrm.Insert(data)
		if err == nil {
			if id > 0 {
				txOrm.QueryTable(&AdminCount{}).Filter("id", 1).Update(orm.Params{
					"count": orm.ColValue(orm.ColAdd, 1),
				})
			} else {
				return errors.New("添加失败")
			}
		}
		return err
	})
	return err == nil

	// db := orm.NewOrm() //创建新orm对象
	// id, err := db.Insert(&data)
	// if id != 0 {
	// 	if err == nil {
	// 		fmt.Println(id)
	// 		return true
	// 	}
	// }
	// return false
}

func AdminUpdate(data Admin) bool {
	db := orm.NewOrm()
	id, err := db.Update(&data)
	if id != 0 {
		if err == nil {
			fmt.Println(id)
			return true
		}
	}
	return false
}

func AdminDelete(id int64) bool {
	db := orm.NewOrm()
	//闭包事务
	err := db.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// data
		num, err := txOrm.Delete(&Admin{Id: id})
		if err == nil {
			if num > 0 {
				txOrm.QueryTable(&AdminCount{}).Filter("id", 1).Update(orm.Params{
					"count": orm.ColValue(orm.ColMinus, 1),
				})
			} else {
				return errors.New("已删除")
			}
		}
		return err
	})
	return err == nil

	// if num, err := db.Delete(&Admin{Id: id}); err == nil {
	// 	if num > 0 {
	// 		return true
	// 	}
	// }
	// return false
}

func GetAdminRoles(uid int64) ([]Role, error) {
	o := orm.NewOrm()
	var roles []Role
	_, err := o.Raw("SELECT r.id,r.group,r.title,r.auth FROM biny_admin_role ar LEFT JOIN `biny_role` r on ar.`role_id` = r.id where ar.admin_id = ? ", uid).QueryRows(&roles)
	return roles, err
}
