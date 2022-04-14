package sms

import (
	"encoding/json"
	"sync"

	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
	"github.com/diy0663/go_project_packages/logger"
)

var once sync.Once

type AliyunSms struct {
	GatewayUrl      string
	AccessKeyId     string
	AccessKeySecret string
	SignName        string
}

type Message struct {
	// 阿里云短信模板编码
	Template string
	// 需要替换的模板中的自定义数据放这里
	Data map[string]string
	// 对上面的Template 进行备注描述而已
	Content string
}

var internalAliyunSMS *AliyunSms

func NewAliyunSms() *AliyunSms {
	once.Do(func() {
		internalAliyunSMS = &AliyunSms{}
	})
	return internalAliyunSMS
}

func (sms *AliyunSms) Send(phone string, message Message) bool {
	smsClient := aliyunsmsclient.New(sms.GatewayUrl)
	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析绑定错误", err.Error())
		return false
	}

	result, err := smsClient.Execute(
		sms.AccessKeyId,
		sms.AccessKeySecret,
		phone,
		sms.SignName,
		message.Template,
		string(templateParam),
	)

	logger.DebugJSON("短信[阿里云]", "请求内容", smsClient.Request)
	logger.DebugJSON("短信[阿里云]", "接口响应", result)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "发信失败", err.Error())
		return false
	}
	resultJSON, err := json.Marshal(result)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析响应 JSON 错误", err.Error())
		return false
	}

	if result.IsSuccessful() {
		logger.DebugString("短信[阿里云]", "发信成功", "")
		return true
	} else {
		logger.ErrorString("短信[阿里云]", "服务商返回错误", string(resultJSON))
		return false
	}

}
