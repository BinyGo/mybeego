/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-22 16:11:54
 * @LastEditTime: 2022-01-22 18:28:47
 */
package validate

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/validation"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

type LoginValidate struct {
	Username string `json:"username" valid:"Required"`
	Password string `json:"password" valid:"Required;MinSize(2)"`
	Mobile   string `json:"mobile"`
}

func ValidLogin(login *LoginValidate) error {

	valid := validation.Validation{}
	valid.Required(login.Username, "用户名").Message("不能为空")
	valid.Required(login.Password, "密码").Message("不能为空")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		message := ""
		for _, err := range valid.Errors {
			//fmt.Println(err.Key, err.Message)
			message = err.Key + err.Message
		}
		return errors.New(message)
	}
	return nil
}

func ValidLogin2(login *LoginValidate) error {
	validation.SetDefaultMessage(MessageTmpls)
	valid := validation.Validation{}

	// 5.验证数据
	b, err := valid.Valid(login) // err是指结构体的定义有没有问题，b是指校验有没有问题
	if err != nil {              // 判断结构体定义有没有错误
		fmt.Println("struct error,", err)
		return err
	}
	if !b { // 判断数据校验有没有错误
		message := ""
		for _, dataerr := range valid.Errors {
			message = message + dataerr.Message
			fmt.Printf("%s: %s\n", dataerr.Key, dataerr.Message)
		}
		return errors.New(message)
	}
	return nil
}

var MessageTmpls = map[string]string{
	"Required":     "不能为空",
	"Min":          "最小值为 %d",
	"Max":          "最大值为 %d",
	"Range":        "Range is %d to %d",
	"MinSize":      "Minimum size is %d",
	"MaxSize":      "Maximum size is %d",
	"Length":       "Required length is %d",
	"Alpha":        "Must be valid alpha characters",
	"Numeric":      "Must be valid numeric characters",
	"AlphaNumeric": "Must be valid alpha or numeric characters",
	"Match":        "Must match %s",
	"NoMatch":      "Must not match %s",
	"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
	"Email":        "Must be a valid email address",
	"IP":           "Must be a valid ip address",
	"Base64":       "Must be valid base64 characters",
	"Mobile":       "Must be valid mobile number",
	"Tel":          "Must be valid telephone number",
	"Phone":        "Must be valid telephone or mobile phone number",
	"ZipCode":      "Must be valid zipcode",
}
