package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
	"server-bof/config"
	"server-bof/database"
	db "server-bof/db/sqlc"
	"strings"
)

func AddReview(c *fiber.Ctx) error {
	secretKey := config.GetEnv("JWT_SECRET")
	cookie := c.Cookies("jwtToken")
	id := c.Params("product_id")
	store := db.NewStore(database.DB)

	dataReview := new(db.Review)
	if err := c.BodyParser(&dataReview); err != nil {
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

	arg := db.CreateReviewParams{
		ReviewID:         strings.ToUpper(IdGeneration()),
		ProductReviewID:  id,
		CustomerReviewID: claims.Issuer,
		Review:           dataReview.Review,
	}

	_, err = store.CreateReview(context.Background(), arg)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Review added successfully",
	})
}

func UpdateReview(c *fiber.Ctx) error {

	store := db.NewStore(database.DB)
	secretKey := config.GetEnv("JWT_SECRET")
	cookie := c.Cookies("jwtToken")

	data := new(db.Review)
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
	customer, err := store.GetCustomerWithId(context.Background(), claims.Issuer)
	if err != nil {
		return err
	}

	if customer.CustomerID != claims.Issuer {
		return nil
	}

	arg := db.UpdateReviewParams{
		Review:   data.Review,
		ReviewID: data.ReviewID,
	}
	err = store.UpdateReview(context.Background(), arg)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "review updated successfully",
	})
}

func RemoveReview(c *fiber.Ctx) error {

	secretKey := config.GetEnv("JWT_SECRET")
	cookie := c.Cookies("jwtToken")
	store := db.NewStore(database.DB)

	data := new(db.Review)
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

	err = store.DeleteReview(context.Background(), data.ReviewID)
	if err != nil {
		return err
	}
	return c.JSON("Review deleted successfully")
}
