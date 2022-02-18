/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-19 21:36:37
 * @LastEditTime: 2022-01-19 23:41:51
 */
package initial

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"

	"github.com/beego/beego/v2/client/orm"
)

func initDatabase() {
	mysql_user, _ := beego.AppConfig.String("mysql_user")
	mysql_pass, _ := beego.AppConfig.String("mysql_pass")
	mysql_host, _ := beego.AppConfig.String("mysql_host")
	mysql_port, _ := beego.AppConfig.Int("mysql_port")
	mysql_db_name, _ := beego.AppConfig.String("mysql_db_name")

	//fmt.Sprintf("%s:%s(%s:%d)@/%s?charset=utf8\n", mysql_user, mysql_pass, mysql_host, mysql_port, mysql_db_name)

	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.31.22:3306)/tp6-admin?charset=utf8")
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", mysql_user, mysql_pass, mysql_host, mysql_port, mysql_db_name))

}
