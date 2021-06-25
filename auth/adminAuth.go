package auth

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	"server-bof/config"
	"server-bof/database"
	"server-bof/models"
)

func IsAdminAuth(c *fiber.Ctx) bool {
	secretKey := config.GetEnv("JWT_SECRET")

	cookie := c.Cookies("jwtToken")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusUnauthorized)
		return false
	}
	claims := token.Claims.(*jwt.StandardClaims)

	res, err := database.DB.Query("SELECT  * FROM users WHERE user_id = ?", claims.Issuer)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		log.Println(err)
		return false
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)
	var user models.User
	if res.Next() {
		err := res.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Password)
		if err != nil {
			return false
		}
	}
	if user.Role == "SuperUser" {
		return true
	}
	return false
}
