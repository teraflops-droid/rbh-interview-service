package common

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/teraflops-droid/rbh-interview-service/configuration"
	"github.com/teraflops-droid/rbh-interview-service/exception"
	"strconv"
	"time"
)

func GenerateToken(username string, role string, config configuration.Config) string {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	claims := jwt.MapClaims{
		"username": username,
		"roles":    role,
		"exp":      time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}
