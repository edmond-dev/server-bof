package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"log"
	"strings"
	"testing"
)

//Creating a test product
func createTestProduct(t *testing.T) Product {

	category := createCategoryTest(t)

	// list all products
	arg := CreateProductParams{
		ProductID:          strings.ToUpper(uuid.New().String()),
		CategoryID:         category.CategoryID,
		ImageUrlPublicID:   "some public product image url",
		ImageUrlSecureID:   "some product image url",
		ProductName:        "some product name",
		ProductDescription: "some product description",
		Price:              95,
		QuantityInStock:    "77",
	}
	_, err := testQueries.CreateProduct(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, category)

	product, err := testQueries.GetProduct(context.Background(), arg.ProductID)
	if err != nil {
		log.Fatal(err.Error())
	}
	require.NoError(t, err)
	require.NotEmpty(t, product)

	return product
}

//Testing creating a product
func TestCreateProduct(t *testing.T) {
	createTestProduct(t)

}

//Test getting products form database
func TestGetProducts(t *testing.T) {
	products, err := testQueries.ListProducts(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, products)
}

//Testing getting categories and products
func TestGetCategoriesAndProducts(t *testing.T) {
	category := createTestProduct(t)

	categoryAndProducts, err := testQueries.GetCategoryAndProducts(context.Background(), category.CategoryID)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.NotEmpty(t, categoryAndProducts)

}

//Testing updating a product
func TestUpdateProduct(t *testing.T) {
	product := createTestProduct(t)

	arg := UpdateProductParams{
		ImageUrlPublicID:   product.ImageUrlPublicID,
		ImageUrlSecureID:   product.ImageUrlSecureID,
		ProductName:        product.ProductName,
		ProductDescription: "updated",
		Price:              product.Price,
		QuantityInStock:    product.QuantityInStock,
		ProductID:          product.ProductID,
	}

	err := testQueries.UpdateProduct(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, product)
}

//Test deleting a product from the database
func TestDeleteProduct(t *testing.T) {
	product := createTestProduct(t)

	err := testQueries.DeleteProduct(context.Background(), product.ProductID)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, product)
}
