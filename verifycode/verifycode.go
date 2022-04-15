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
				RedisClient: &redis.RedisClient{},
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
