// Created by Hisen at 2019-07-08.
package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
	"net/http"
	"time"
)

var jwtSecret []byte

type LoginUser struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func idGen() (id uint64, err error) {
	setting := sonyflake.Settings{}
	sf := sonyflake.NewSonyflake(setting)
	return sf.NextID()
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
func validateTokenmiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("auth")
		_, err := ParseToken(token)
		code := 0
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = 1
			default:
				code = 2

			}
		}
		if code != 0 {
			context.JSON(http.StatusOK, gin.H{"msg": "验证失败", "code": code})
			context.Abort()
			return
		}
		context.Next()

	}

}

func tokenGen(username, password string) (tokenString string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(30 * time.Minute)

	claims := Claims{
		EncodeMD5(username),
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret = []byte("hello world")
	tokenString, err = tokenClaims.SignedString(jwtSecret)
	return
}

func loginHandler(c *gin.Context) {
	var userinfo LoginUser
	if e := c.ShouldBind(&userinfo); e != nil {
		c.JSON(http.StatusOK, gin.H{"status": e.Error()})
		return
	} else {
		username := userinfo.Username
		password := userinfo.Password
		if username == "hanxin" && password == "123" {
			tokenString, err := tokenGen(username, password)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": tokenString})

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"code": "no auth"})
			return
		}
	}

}
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 11})
}
func main() {
	app := gin.Default()

	app.POST("/login", loginHandler)
	app.GET("/index", validateTokenmiddleware(), indexHandler)
	app.Run(":9988")
}
