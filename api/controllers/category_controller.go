package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"server-bof/auth"
	"server-bof/database"
	db "server-bof/db/sqlc"
	"strings"
)

func AddCategory(c *fiber.Ctx) error {

	auth.IsAdminAuth(c)

	store := db.NewStore(database.DB)
	data := new(db.Category)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	arg := db.CreateCategoryParams{
		CategoryID:              strings.ToUpper(IdGeneration()),
		CategoryName:            data.CategoryName,
		CategoryHtmlDescription: data.CategoryHtmlDescription,
		Image:                   data.Image,
	}

	_, err := store.CreateCategory(context.Background(), arg)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("Category added successfully!")

}

func GetCategories(c *fiber.Ctx) error {
	store := db.NewStore(database.DB)

	categories, err := store.GetCategories(context.Background())
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(categories)
}

func GetCategory(c *fiber.Ctx) error {

	auth.IsAdminAuth(c)

	categoryId := c.Params("category_id")
	store := db.NewStore(database.DB)

	category, err := store.GetCategory(context.Background(), categoryId)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {

	auth.IsAdminAuth(c)

	id := c.Params("category_id")
	store := db.NewStore(database.DB)
	data := new(db.Category)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	arg := db.UpdateCategoryParams{
		CategoryName:            data.CategoryName,
		CategoryHtmlDescription: data.CategoryHtmlDescription,
		Image:                   data.Image,
		CategoryID:              id,
	}

	err := store.UpdateCategory(context.Background(), arg)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("Category updated successfully!")

}

func DeleteCategory(c *fiber.Ctx) error {

	auth.IsAdminAuth(c)

	id := c.Params("category_id")
	store := db.NewStore(database.DB)

	err := store.DeleteCategory(context.Background(), id)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("Category successfully deleted.")
}
