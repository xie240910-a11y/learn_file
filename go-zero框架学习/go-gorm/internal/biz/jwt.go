package biz

import "github.com/golang-jwt/jwt/v4"

/*
参数1 秘钥
2：当前时间戳
3：过期时间
4:负载信息
*/
func GetJwtToken(srcretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(srcretKey))
}
