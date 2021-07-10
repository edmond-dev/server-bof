package db

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"log"
	"testing"
)

//Testing getting statics
func TestGetTotalProductsPerCategory(t *testing.T) {
	product := createTestProduct(t)

	value, err := testQueries.GetProductStatsPerCategory(context.Background(), product.CategoryID)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.NotEmpty(t, value)

	log.Println(value)
}

//Testing the calculation of the total number of products available
func TestTotalProductsAvailable(t *testing.T) {

	value, err := testQueries.ProductsAvailable(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, value)

	log.Println(value.TotalProductsAvailable)

}

//Testing getting total revenue value
//func TestGetTotalRevenue(t *testing.T)  {
//	value, err := testQueries.GetTotalRevenue(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//	require.NoError(t, err)
//	require.NotEmpty(t, value)
//
//	log.Println(value.TotalRevenue)
//}

//Testing the calculation of the total number of products available
func TestTotalRegisteredCustomer(t *testing.T) {

	value, err := testQueries.GetTotalCustomers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, value)

	log.Println(value.TotalNumberOfCustomers)

}
