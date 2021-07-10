package db

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"log"
	"testing"
)

//Creating a test product
func createTestAddress(t *testing.T) Address {
	customer := createRandomCustomer(t)
	arg := CreateAddressParams{
		AddressCustomerID: customer.CustomerID,
		Email:             customer.Email,
		Street:            "Some street number",
		AddressLine1:      "Some address line 1",
		AddressLine2:      "some address line 2",
		Phone:             "20893497349274",
		City:              "Caldwell",
		State:             "Idaho",
		ZipCode:           "83605",
	}
	_, err := testQueries.CreateAddress(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, customer)

	address, err := testQueries.GetAddress(context.Background(), arg.AddressCustomerID)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, address)

	return address
}

//Testing creating address
func TestCreateAddress(t *testing.T) {
	createTestAddress(t)
}

//Testing update address
func TestUpdateAddress(t *testing.T) {
	address := createTestAddress(t)

	arg := UpdateAddressParams{
		AddressCustomerID: address.AddressCustomerID,
		Email:             address.Email,
		Street:            address.Street,
		AddressLine1:      "Updated",
		AddressLine2:      address.AddressLine2,
		Phone:             address.Phone,
		City:              address.City,
		State:             address.State,
		ZipCode:           address.ZipCode,
	}

	err := testQueries.UpdateAddress(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, address)
}

//Testing delete address
func TestDeleteAddress(t *testing.T) {
	address := createTestAddress(t)

	err := testQueries.DeleteAddress(context.Background(), address.AddressCustomerID)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, address)
}
