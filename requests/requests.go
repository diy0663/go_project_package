package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

//  针对 Gin 做 JSON 解析和验证器封装
// 使用 Govalidator
// go get github.com/thedevsaddam/govalidator

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

// 被API调用
func ValidateInAPI(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	//ShouldBind方法可以处理不同格式的数据，包括JSON、XML、Form表单等，而ShouldBindJSON只能处理JSON格式的数据。
	//ShouldBind方法会根据请求的Content-Type自动选择绑定器，而ShouldBindJSON只能处理Content-Type为application/json的请求。
	// 如果严格要求请求的数据格式是JSON，建议使用ShouldBindJSON方法；
	// 如果请求的数据格式不确定或可能是多种格式，可以使用ShouldBind方法。
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头",
			"error":   err.Error(),
		})

		return false
	}
	errs := handler(obj, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求验证不通过，具体请查看 errors",
			"errors":  errs,
		})
		return false
	}

	return true
}

// 在 ValidatorFunc 里面被使用, 而 ValidatorFunc 是每个request单独验证里面需要用到的方法
func ValidateInRequest(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 验证器配置
	opts := govalidator.Options{
		// 最里面会检查data 是否为指针类型
		Data:     data,
		Rules:    rules,
		Messages: messages,
		// 结构体的 标签标识符
		TagIdentifier: "valid",
	}
	return govalidator.New(opts).ValidateStruct()
}
