package requests

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 表单请求数据, 里面只有一个email 参数
type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// 具体的表单校验函数,定制了报错信息, 函数的参数也可以只被声明，不被使用 , 所以第二个参数 c 没被使用也没关系
func SignupEmailExistFunc(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	return validate(data, rules, messages)
}

type APIController struct {
}

func (api *APIController) CheckEmailExist(c *gin.Context) {
	request := SignupEmailExistRequest{}
	ok := Validate(c, &request, SignupEmailExistFunc)
	if !ok {
		fmt.Println("验证不通过:", ok)
		return
	}
	// 继续你的业务逻辑判断

}
