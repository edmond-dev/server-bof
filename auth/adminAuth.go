package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"server-bof/database"
	db "server-bof/db/sqlc"
)

func IsAdminAuth(c *fiber.Ctx) bool {

	cookie := c.Cookies("jwtToken")
	secretKey := os.Getenv("JWT_SECRET")
	store := db.NewStore(database.DB)

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		log.Fatal(err.Error())
	}
	claims := token.Claims.(*jwt.StandardClaims)

	res, err := store.GetCustomerWithId(context.Background(), claims.Issuer)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		log.Fatal(err.Error())
	}
	if res.Role == "SuperUser" {
		return true
	}
	return false
}
