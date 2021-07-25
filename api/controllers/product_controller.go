package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"server-bof/auth"
	"server-bof/database"
	db "server-bof/db/sqlc"
	"strings"
)

func AddProduct(c *fiber.Ctx) error {

	auth.IsAdminAuth(c)

	store := db.NewStore(database.DB)
	data := new(db.Product)
	category := new(db.Category)

	if err := c.BodyParser(&category); err != nil {
		log.Fatal(err.Error())
	}
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	arg := db.CreateProductParams{
		ProductID:          strings.ToUpper(IdGeneration()),
		CategoryID:         category.CategoryID,
		ImageUrlPublicID:   data.ImageUrlPublicID,
		ImageUrlSecureID:   data.ImageUrlSecureID,
		ProductName:        data.ProductName,
		ProductDescription: data.ProductDescription,
		Price:              data.Price,
		QuantityInStock:    data.QuantityInStock,
	}
	_, err := store.CreateProduct(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(fiber.Map{
		"message": "product added successfully",
	})
}

func GetProducts(c *fiber.Ctx) error {
	store := db.NewStore(database.DB)

	rows, err := store.ListProducts(context.Background())
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(rows)

}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("product_id")
	store := db.NewStore(database.DB)

	rows, err := store.GetProduct(context.Background(), id)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(rows)

}

func UpdateProduct(c *fiber.Ctx) error {

	auth.IsAdminAuth(c)
	//id := c.Params("product_id")
	store := db.NewStore(database.DB)

	data := new(db.Product)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	arg := db.UpdateProductParams{
		ImageUrlPublicID:   data.ImageUrlPublicID,
		ImageUrlSecureID:   data.ImageUrlSecureID,
		ProductName:        data.ProductName,
		ProductDescription: data.ProductDescription,
		Price:              data.Price,
		QuantityInStock:    data.QuantityInStock,
	}
	err := store.UpdateProduct(context.Background(), arg)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func RemoveProduct(c *fiber.Ctx) error {

	auth.IsAdminAuth(c)

	id := c.Params("product_id")
	store := db.NewStore(database.DB)

	data := new(db.Product)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	err := store.DeleteProduct(context.Background(), id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
