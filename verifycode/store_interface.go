package verifycode

// 定义数字验证码(短信或email发送)相关的存储接口规范 (等着被redis或memcache的等实现)
type Store interface {
	// 保存验证码到存储器
	Set(id string, value string) bool
	// 从存储器中获取验证码的值, 第二个参数代表是否取完就清除存储值
	Get(id string, is_clear bool) string
	// 验证验证码的正确性
	Verify(id string, answer string, is_clear bool) bool
}
