package verifycode

import (
	"sync"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/helper"
	"github.com/diy0663/go_project_packages/redis"
)

type VerifyCode struct {
	StoreType Store
}

var once sync.Once

var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {

	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			StoreType: &RedisStore{
				// 这里实例化的时候没有传入任何redis配置信息是因为使用了redis全局变量,这个变量在项目启动的时候需要一开始就被初始化
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

func (vc *VerifyCode) Generate(key string) string {
	code := helper.RandomNumber(config.GetInt("verifycode.code_length"))
	vc.StoreType.Set(key, code)
	return code
}

func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {
	return vc.StoreType.Verify(key, answer, false)
}
