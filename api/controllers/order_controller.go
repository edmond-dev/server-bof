package controllers

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"server-bof/database"
	db "server-bof/db/sqlc"
	"strings"
)

func AddOrder(c *fiber.Ctx) error {
	store := db.NewStore(database.DB)
	data := new(db.Order)
	product := new(db.Product)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwtToken")
	secretKey := os.Getenv("JWT_SECRET")

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

	arg := db.CreateOrderParams{
		OrderNumber:       strings.ToUpper(IdGeneration()),
		CustomerID:        claims.Issuer,
		ProductID:         product.ProductID,
		OrderCategoryName: data.OrderCategoryName,
		QuantityOrdered:   data.QuantityOrdered,
		PriceEach:         data.PriceEach,
		CustomerComments:  data.CustomerComments,
	}

	_, err = store.CreateOrder(context.Background(), arg)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("Order Placed successfully!")
}

func GetOrder(c *fiber.Ctx) error {
	store := db.NewStore(database.DB)
	orderId := c.Params("order_id")

	order, err := store.GetOrder(context.Background(), orderId)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	orderId := c.Params("order_id")
	store := db.NewStore(database.DB)

	err := store.DeleteOrder(context.Background(), orderId)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON("OrderSuccessfully deleted")
}
