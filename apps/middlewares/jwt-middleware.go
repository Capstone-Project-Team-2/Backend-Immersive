package middlewares

import (
	"capstone-tickets/apps/config"
	"time"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.JWT_SECRRET),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId string, userRole string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["userRole"] = userRole
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_SECRRET))
}

func ExtractToken(e echo.Context) (string, string, string) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)
		role := claims["role"].(string)
		email := claims["emails"].(string)
		return userId, role, email
	}
	return "", "", ""
}
