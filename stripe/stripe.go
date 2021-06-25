package stripe

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"server-bof/config"
	"server-bof/database"
	"server-bof/models"
)

func Payment(c *fiber.Ctx) error {

	stripeApiKey := config.GetEnv("STRIPE_APIKEY")
	secretKey := config.GetEnv("JWT_SECRET")

	stripe.Key = stripeApiKey
	cookie := c.Cookies("jwtToken")

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

	user, err := database.DB.Query("SELECT user_id, email FROM users WHERE user_id = ?", claims.Issuer)
	if err != nil {
		return err
	}
	defer func(user *sql.Rows) {
		err := user.Close()
		if err != nil {

		}
	}(user)

	userInfo := models.User{}

	for user.Next() {
		if err := user.Scan(&userInfo.UserID, &userInfo.Email); err != nil {
			return err
		}
	}

	data := new(models.StripePayment)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	response, err := paymentintent.New(&stripe.PaymentIntentParams{
		Amount:       stripe.Int64(data.Amount),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String(data.ProductName),
		ReceiptEmail: stripe.String(userInfo.Email),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	})

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	res, err := database.DB.Query(`
		INSERT INTO orders (order_id, user_id, amount, productName, receiptEmail, street, address_line_1, address_line_2, City, state, zip_code) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ? , ?)`,
		xid.New(), userInfo.UserID, data.Amount, data.ProductName, userInfo.Email, data.Street, data.AddressLine1, data.AddressLine2, data.City, data.State, data.ZipCode)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)

	return c.JSON(response)
}
