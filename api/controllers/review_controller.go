package controllers

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"server-bof/config"
	"server-bof/database"
	"server-bof/models"
)

func AddReview(c *fiber.Ctx) error {

	secretKey := config.GetEnv("JWT_SECRET")
	cookie := c.Cookies("jwtToken")

	data := new(models.Review)
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	product := new(models.Product)
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	res, err := database.DB.Query("SELECT * FROM users WHERE user_id = ?", claims.Issuer)
	if err != nil {
		return err
	}

	user := models.User{}
	if res.Next() {

		err := res.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Password)
		if err != nil {
			return err
		}
	}
	if user.Email == "" {
		return c.JSON(fiber.Map{
			"error": "Failed to add review",
		})
	}
	result, err := database.DB.Query(`
		INSERT INTO reviews (review_id, user_id,product_id, reviews, name, email)
		VALUES (?, ?, ?, ?, ?, ?)`,
		xid.New(), user.UserID, product.ProductID, data.Reviews, user.FirstName, user.Email)

	if err != nil {
		return err
	}
	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {

		}
	}(result)

	return c.JSON(fiber.Map{
		"message": "Review added successfully",
	})
}

func GetReview(c *fiber.Ctx) error {
	review := models.Review{}

	data := new(models.Product)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	rows, err := database.DB.Query(`
		SELECT r.review_id, r.product_id, r.user_id, r.reviews, r.name, r.email
		FROM products p
		JOIN reviews r on p.product_id = r.product_id
		WHERE p.product_id = ?
		`, data.ProductID)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {

		if err := rows.Scan(
			&review.ReviewID,
			&review.ProductID,
			&review.UserID,
			&review.Reviews,
			&review.Name,
			&review.Email,
		); err != nil {
			return err
		}

	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	return c.JSON(review)

}

func UpdateReview(c *fiber.Ctx) error {

	secretKey := config.GetEnv("JWT_SECRET")
	cookie := c.Cookies("jwtToken")

	data := new(models.Review)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	_, err = database.DB.Query("SELECT * FROM users WHERE user_id = ?", claims.Issuer)
	if err != nil {
		return err
	}
	res, err := database.DB.Query("UPDATE reviews SET reviews = ? WHERE review_id = ?", data.Reviews, data.ReviewID)
	if err != nil {
		return err
	}

	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)

	return c.JSON(fiber.Map{
		"message": "review updated successfully",
	})
}

func RemoveReview(c *fiber.Ctx) error {

	secretKey := config.GetEnv("JWT_SECRET")
	cookie := c.Cookies("jwtToken")

	data := new(models.Review)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthenticated",
		})
	}

	res, err := database.DB.Query("DELETE FROM reviews WHERE review_id = ?", data.ReviewID)
	if err != nil {
		return err
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)

	return c.JSON("Review deleted successfully")
}