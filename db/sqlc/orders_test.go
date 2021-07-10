package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"server-bof/util"
	"testing"
)

//Testing creating an order
func createTestOrder(t *testing.T) Order {
	product := createTestProduct(t)
	category := createCategoryTest(t)
	customer := createRandomCustomer(t)

	arg := CreateOrderParams{
		OrderNumber:       util.RandomString(7),
		CustomerID:        customer.CustomerID,
		ProductID:         product.ProductID,
		OrderCategoryName: category.CategoryName,
		QuantityOrdered:   util.RandomInt(1, 70),
		PriceEach:         util.RandomInt(1, 500),
		CustomerComments:  util.RandomString(20),
	}

	_, err := testQueries.CreateOrder(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)

	order, err := testQueries.GetOrder(context.Background(), arg.OrderNumber)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, order)

	return order

}

//Testing creating an order
func TestCreateOrder(t *testing.T) {
	createTestOrder(t)
}

//Testing updating an order
func TestUpdateOrder(t *testing.T) {
	order := createTestOrder(t)

	arg := UpdateOrderParams{
		CustomerComments: util.RandomString(20),
		OrderNumber:      order.OrderNumber,
	}

	err := testQueries.UpdateOrder(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, order)
}

//Testing updating an order
func TestDeleteOrder(t *testing.T) {
	order := createTestOrder(t)

	err := testQueries.DeleteOrder(context.Background(), order.OrderNumber)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, order)
}
