package verifycode

import (
	"time"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/redis"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(id string, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_minutes"))
	if config.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_minutes"))
	}

	return s.RedisClient.Set(s.KeyPrefix+id, value, ExpireTime)
}

func (s *RedisStore) Get(id string, is_clear bool) string {
	key := s.KeyPrefix + id
	value := s.RedisClient.Get(key)
	if is_clear {
		s.RedisClient.Del(key)
	}
	return value

}

// 验证验证码的正确性
func (s *RedisStore) Verify(id string, answer string, is_clear bool) bool {
	value := s.Get(id, is_clear)
	return value == answer
}
