package user_ser

import (
	"Goblog/service/redis_ser"
	"Goblog/utils/jwts"
	"time"
)

type UserService struct {
}

// Logout 是否登出
func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	//过期时间 需要计算过期时间 距离现在的过期时间
	exp := claims.ExpiresAt
	now := time.Now()
	// 截至时间
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)
}
