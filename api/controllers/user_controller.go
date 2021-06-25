package controllers

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"server-bof/config"
	"server-bof/database"
	"server-bof/models"
	"time"
)

func UserCtrlRegister(c *fiber.Ctx) error {
	data := new(models.User)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword(data.Password, 10)

	res, err := database.DB.Query(`INSERT INTO users (user_id, first_name, last_name, email, role, password) VALUES (?, ?, ?, ?, ?, ?)`, xid.New(), data.FirstName, data.LastName, data.Email, "user", password)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)

	return c.JSON("Registration successfully!")
}

////user log in method

func UserCtrlLogin(c *fiber.Ctx) error {
	secretKey := config.GetEnv("JWT_SECRET")
	var user models.User

	data := new(models.User)
	if err := c.BodyParser(data); err != nil {
		return err
	}
	res, err := database.DB.Query("SELECT * FROM users WHERE email = ?", data.Email)
	if err != nil {
		return err
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)
	if res.Next() {
		err := res.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Password)
		if err != nil {
			return err
		}
	} else {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"Error": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, data.Password); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "Incorrect email or password",
		})
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.UserID,
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
	secretKey := config.GetEnv("JWT_SECRET")

	var user models.User

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

	res, err := database.DB.Query("SELECT * FROM users WHERE user_id = ?", claims.Issuer)
	if err != nil {
		return err
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)

	if res.Next() {
		err := res.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Password)
		if err != nil {
			return err
		}
	} else {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"Error": "User not found",
		})
	}

	return c.JSON(user)
}

//Update user preferences

func UpdateUser(c *fiber.Ctx) error {
	secretKey := config.GetEnv("JWT_SECRET")

	data := new(models.User)
	er := c.BodyParser(&data)
	if er != nil {
		return er
	}
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

	res, err := database.DB.Query(`
		UPDATE users SET first_name = ?, last_name = ?, email = ?
		WHERE user_id = ?`,
		data.FirstName, data.LastName, data.Email, claims.Issuer)
	if err != nil {
		return err
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)
	return c.JSON("updated successfully!")
}

func CreateAdmin(c *fiber.Ctx) error {
	var user models.User

	data := new(models.User)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	res, err := database.DB.Query("SELECT * FROM users WHERE email = ?", data.Email)
	if err != nil {
		return err
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)
	if res.Next() {
		err := res.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Password)
		if err != nil {
			return err
		}
	} else {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"Error": "User not found",
		})
	}

	res, err = database.DB.Query("UPDATE users SET role = ? WHERE user_id = ?", "SuperUser", user.UserID)
	if err != nil {
		return err
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)

	return c.JSON(fiber.Map{
		"message": "Super user created successfully!",
	})
}