package controllers

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"os"
	"server-bof/config"
	"server-bof/database"
	db "server-bof/db/sqlc"
	"strings"
	"time"
)

func UserCtrlRegister(c *fiber.Ctx) error {
	store := db.NewStore(database.DB)

	data := new(db.Customer)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	arg := db.CreateCustomerParams{
		CustomerID: strings.ToUpper(IdGeneration()),
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Email:      data.Email,
		Role:       "customer",
		Password:   string(password),
	}

	_, err := store.CreateCustomer(context.Background(), arg)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("Registration successfully!")
}

//Create Admin

func CreateAdmin(c *fiber.Ctx) error {
	store := db.NewStore(database.DB)

	data := new(db.Customer)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	arg := db.CreateCustomerParams{
		CustomerID: strings.ToUpper(IdGeneration()),
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Email:      data.Email,
		Role:       "SuperUser",
		Password:   string(password),
	}

	_, err := store.CreateCustomer(context.Background(), arg)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("Super user created successfully!")
}

////user log in method

func UserCtrlLogin(c *fiber.Ctx) error {
	store := db.NewStore(database.DB)
	secretKey := config.GetEnv("JWT_SECRET")

	data := new(db.Customer)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	res, err := store.GetCustomerWithEmail(context.Background(), data.Email)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"Error": "User not found",
			"data":  data,
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "Incorrect email or password",
		})
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    res.CustomerID,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	jwtToken, err := jwtClaims.SignedString([]byte(secretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": "Could not login",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwtToken",
		Value:    jwtToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logged In",
	})
}

// user logout method

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwtToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logged Out",
	})
}

func GetUser(c *fiber.Ctx) error {

	cookie := c.Cookies("jwtToken")
	secretKey := os.Getenv("JWT_SECRET")
	store := db.NewStore(database.DB)

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

	res, err := store.GetCustomerWithId(context.Background(), claims.Issuer)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"Error": "User not found",
		})

	}
	return c.JSON(res)
}

func GetCustomerOrderDetails(c *fiber.Ctx) error {
	cookie := c.Cookies("jwtToken")
	secretKey := os.Getenv("JWT_SECRET")
	store := db.NewStore(database.DB)

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

	customerOrderDetails, err := store.GetCustomerOrderDetailsAndAddr(context.Background(), claims.Issuer)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(customerOrderDetails)
}

//Update user preferences

func UpdateUser(c *fiber.Ctx) error {
	store := db.NewStore(database.DB)
	secretKey := os.Getenv("JWT_SECRET")
	cookie := c.Cookies("jwtToken")

	data := new(db.Customer)
	er := c.BodyParser(&data)
	if er != nil {
		return er
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

	arg := db.UpdateCustomerParams{
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Email:      data.Email,
		CustomerID: claims.Issuer,
	}

	err = store.UpdateCustomer(context.Background(), arg)
	if err != nil {
		return err
	}
	return c.JSON("updated successfully!")
}
