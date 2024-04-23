package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"penjadwalan-sidang/internal/model/entity"
	"strings"
)

type AuthConfig struct {
	Filter       func(*fiber.Ctx) error
	Unauthorized fiber.Handler
}

func UserAuthentication(c AuthConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		header := ctx.Get("Authorization")
		if header == "" {
			return c.Unauthorized(ctx)
		}

		tokenString := strings.Replace(header, "Bearer ", "", 1)

		authToken := entity.AuthToken{}

		validateJWT, err := ValidateJWT(tokenString)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "JWT NOT VALID",
			})
		}

		user_id := validateJWT["id"].(float64)
		rolesInterface, ok := validateJWT["role"].([]interface{})
		if !ok {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Role conversion error",
			})
		}

		authToken.UserId = int(user_id)
		authToken.Nama = validateJWT["nama"].(string)
		authToken.Username = validateJWT["username"].(string)
		for _, role := range rolesInterface {
			authToken.Role = append(authToken.Role, role.(string))
		}

		ctx.Locals("user_id", authToken.UserId)
		ctx.Locals("role", authToken.Role)
		ctx.Locals("username", authToken.Username)
		ctx.Locals("name", authToken.Nama)

		return ctx.Next()
	}
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	return checkJWT(tokenString, os.Getenv("JWT_KEY"))
}

func checkJWT(tokenString string, secret string) (jwt.MapClaims, error) {
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SIGNATURE_KEY = []byte(secret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
