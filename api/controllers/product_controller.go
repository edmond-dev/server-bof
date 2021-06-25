package controllers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"log"
	"server-bof/auth"
	"server-bof/database"
	"server-bof/models"
)

func AddProduct(c *fiber.Ctx) error {
	if auth.IsAdminAuth(c) {

		data := new(models.Product)
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		_, err := database.DB.Query(`
			INSERT INTO products (product_id, image_url_public_id, image_url_secure_id, product_name, product_description, price)
			VALUES (?, ?, ?, ?, ?, ?)`,
			xid.New(), data.ImageUrlPublicID, data.ImageUrlSecureID, data.ProductName, data.ProductDescription, data.Price)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(fiber.Map{
			"message": "product added successfully",
		})
	}
	return c.JSON(fiber.Map{
		"error": "Failed to add product. Not an Admin",
	})
}

func GetProducts(c *fiber.Ctx) error {
	results := models.Products{}
	product := models.Product{}
	rows, err := database.DB.Query(`SELECT *  FROM products`)

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
			&product.ProductID,
			&product.ImageUrlPublicID,
			&product.ImageUrlSecureID,
			&product.ProductName,
			&product.ProductDescription,
			&product.Price,
		); err != nil {
			return err
		}
		results.Products = append(results.Products, product)

	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	return c.JSON(results)

}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("product_id")

	product := models.Product{}

	data := new(models.Product)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	rows, err := database.DB.Query(`
		SELECT *
		FROM products p
		WHERE p.product_id = ?
		`, id)

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
			&product.ProductID,
			&product.ImageUrlPublicID,
			&product.ImageUrlSecureID,
			&product.ProductName,
			&product.ProductDescription,
			&product.Price,
		); err != nil {
			return err
		}

	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	return c.JSON(product)

}

func UpdateProduct(c *fiber.Ctx) error {

	id := c.Params("product_id")

	data := new(models.Product)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if auth.IsAdminAuth(c) {
		res, err := database.DB.Query(`
			UPDATE products
			SET image_url_public_id = ?, image_url_secure_id = ?, product_name = ?, product_description = ?, price = ?
			WHERE product_id = ?`,
			data.ImageUrlSecureID, data.ImageUrlPublicID, data.ProductName, data.ProductDescription, data.Price, id)
		if err != nil {
			return err
		}
		defer func(res *sql.Rows) {
			err := res.Close()
			if err != nil {

			}
		}(res)
		return c.JSON(fiber.Map{
			"message": "Product updated successfully",
		})
	}

	return c.JSON(fiber.Map{
		"error": "Product update failed. Not an Admin",
	})
}

func RemoveProduct(c *fiber.Ctx) error {

	id := c.Params("product_id")

	data := new(models.Product)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if auth.IsAdminAuth(c) {
		res, err := database.DB.Query("DELETE FROM products WHERE product_id = ?", id)
		if err != nil {
			return err
		}
		defer func(res *sql.Rows) {
			err := res.Close()
			if err != nil {

			}
		}(res)
		return c.JSON(fiber.Map{
			"message": "Product deleted successfully",
		})
	}

	return c.JSON(fiber.Map{
		"error": "Could not delete product. Not an Admin",
	})
}

