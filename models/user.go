/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-19 15:30:01
 * @LastEditTime: 2022-01-20 10:08:19
 */
package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"` //`json:"-" orm:"-" orm:"column(id)`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	Group      int    `json:"group"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func init() {
	orm.RegisterModelWithPrefix("biny_", new(User))
}

func GetUserByUsername(username string) (User, error) {
	db := orm.NewOrm()
	admin := User{Username: username}
	err := db.Read(&admin, "Username")
	return admin, err
}

func GetUserById(id int64) (User, error) {
	db := orm.NewOrm()
	admin := User{Id: id}
	err := db.Read(&admin)
	return admin, err
}

func GetUserByList(limit int, offset int) ([]*User, error) {
	o := orm.NewOrm()
	//qs = o.QueryTable(&User)
	var admins []*User
	admin := new(User)
	_, err := o.QueryTable(admin).Limit(limit, offset).All(&admins) //All(&admins, "Id", "Username")
	//fmt.Println(qs)
	return admins, err
}

// func GetUserCount() (int64, error) {
// 	db := orm.NewOrm()
// 	userCount := UserCount{Id: 1}
// 	err := db.Read(&userCount)
// 	return userCount.Count, err

// }

func UserCreate(data User) bool {
	db := orm.NewOrm() //创建新orm对象
	id, err := db.Insert(&data)
	if id != 0 {
		if err == nil {
			fmt.Println(id)
			return true
		}
	}
	return false
}

func UserUpdate(data User) bool {
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

func UserDelete(id int64) bool {
	db := orm.NewOrm()
	if num, err := db.Delete(&User{Id: id}); err == nil {
		if num > 0 {
			return true
		}
	}
	return false
}
