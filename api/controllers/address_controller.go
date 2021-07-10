package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"server-bof/database"
	db "server-bof/db/sqlc"
	"strings"
)

func AddAddress(c *fiber.Ctx) error {
	store := db.NewStore(database.DB)
	data := new(db.Address)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	arg := db.CreateAddressParams{
		AddressCustomerID: strings.ToUpper(IdGeneration()),
		Email:             data.Email,
		Street:            data.Street,
		AddressLine1:      data.AddressLine1,
		AddressLine2:      data.AddressLine2,
		Phone:             data.Phone,
		City:              data.City,
		State:             data.State,
		ZipCode:           data.ZipCode,
	}
	_, err := store.CreateAddress(context.Background(), arg)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON("Address added")

}

func UpdateAddress(c *fiber.Ctx) error {
	addressId := c.Params("address_id")
	store := db.NewStore(database.DB)
	data := new(db.Address)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	arg := db.UpdateAddressParams{
		Email:             data.Email,
		Street:            data.Street,
		AddressLine1:      data.AddressLine2,
		AddressLine2:      data.AddressLine1,
		Phone:             data.Phone,
		City:              data.City,
		State:             data.State,
		ZipCode:           data.ZipCode,
		AddressCustomerID: addressId,
	}

	err := store.UpdateAddress(context.Background(), arg)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("Address Updated.")

}

func GetAddress(c *fiber.Ctx) error {
	addressId := c.Path("address_id")
	store := db.NewStore(database.DB)

	address, err := store.GetAddress(context.Background(), addressId)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(address)
}
func DeleteAddress(c *fiber.Ctx) error {
	addressId := c.Params("address_id")
	store := db.NewStore(database.DB)

	err := store.DeleteAddress(context.Background(), addressId)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON("Address Deleted successfully")
}
