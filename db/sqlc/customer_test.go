package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"log"
	"server-bof/util"
	"testing"
)

func createRandomCustomer(t *testing.T) Customer {
	arg := CreateCustomerParams{
		CustomerID: util.RandomString(77),
		FirstName:  util.RandomFirstName(),
		LastName:   util.RandomLastName(),
		Email:      util.RandomEmail(),
		Password:   util.RandomPass(),
	}
	_, err := testQueries.CreateCustomer(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}

	customer, err := testQueries.GetCustomerWithId(context.Background(), arg.CustomerID)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.CustomerID, customer.CustomerID)
	require.Equal(t, arg.FirstName, customer.FirstName)
	require.Equal(t, arg.LastName, customer.LastName)
	require.Equal(t, arg.Email, customer.Email)

	require.NotZero(t, customer.CustomerID)
	require.NotZero(t, customer.CreatedAt)

	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	customer2, err := testQueries.GetCustomerWithId(context.Background(), customer1.CustomerID)
	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.CustomerID, customer2.CustomerID)
	require.Equal(t, customer1.FirstName, customer2.FirstName)
	require.Equal(t, customer1.LastName, customer2.LastName)
	require.Equal(t, customer1.Email, customer1.Email)

	require.NotZero(t, customer2.CustomerID)
	require.NotZero(t, customer1.CreatedAt, customer2.CreatedAt)

}

func TestUpdateCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)

	arg := UpdateCustomerParams{
		CustomerID: customer1.CustomerID,
		FirstName:  customer1.FirstName,
		LastName:   customer1.LastName,
		Email:      customer1.Email,
	}
	err := testQueries.UpdateCustomer(context.Background(), arg)
	customer2, err := testQueries.GetCustomerWithId(context.Background(), customer1.CustomerID)

	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.CustomerID, customer2.CustomerID)
	require.Equal(t, customer1.FirstName, customer2.FirstName)
	require.Equal(t, customer1.LastName, customer2.LastName)
	require.Equal(t, customer1.Email, customer2.Email)

	require.NotZero(t, customer2.CustomerID)
	require.NotZero(t, customer1.CreatedAt, customer2.CreatedAt)

}

func TestDeleteCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)

	err := testQueries.DeleteCustomer(context.Background(), customer1.CustomerID)
	require.NoError(t, err)

	customer2, err := testQueries.GetCustomerWithId(context.Background(), customer1.CustomerID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, customer2)
}
