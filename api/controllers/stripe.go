package controllers

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"os"
	"server-bof/config"
	"server-bof/database"
	db "server-bof/db/sqlc"
)

func Stripe(c *fiber.Ctx) error {

	stripeApiKey := config.GetEnv("STRIPE_APIKEY")
	stripe.Key = stripeApiKey

	store := db.NewStore(database.DB)

	orderData := new(db.Order)
	if err := c.BodyParser(&orderData); err != nil {
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

	customer, err := store.GetCustomerWithId(context.Background(), claims.Issuer)
	if err != nil {
		return c.JSON(err.Error())
	}

	response, err := paymentintent.New(&stripe.PaymentIntentParams{
		Amount:       stripe.Int64(int64(orderData.PriceEach * orderData.QuantityOrdered)),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String(orderData.ProductID),
		ReceiptEmail: stripe.String(customer.Email),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	})

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(response)
}
