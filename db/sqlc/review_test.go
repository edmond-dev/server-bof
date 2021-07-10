package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"log"
	"strings"
	"testing"
)

//Testing creating a review for an existing product
func createTestReview(t *testing.T) Review {
	product := createTestProduct(t)
	customer := createRandomCustomer(t)

	arg := CreateReviewParams{
		ReviewID:         strings.ToUpper(uuid.New().String()),
		ProductReviewID:  product.ProductID,
		CustomerReviewID: customer.CustomerID,
		Review:           "Some customer review of a product",
	}
	_, err := testQueries.CreateReview(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.NotEmpty(t, customer)

	review, err := testQueries.GetReview(context.Background(), arg.ReviewID)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, review)

	return review
}

//Testing creating a review of a product
func TestCreateReview(t *testing.T) {
	createTestReview(t)
}

//Testing updating a product
func TestUpdateReview(t *testing.T) {
	review := createTestReview(t)

	arg := UpdateReviewParams{
		Review:   "some updated review about a product",
		ReviewID: review.ReviewID,
	}

	err := testQueries.UpdateReview(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, review)
}

//Testing deleting a review of a product
func TestDeleteReview(t *testing.T) {
	review := createTestReview(t)

	err := testQueries.DeleteReview(context.Background(), review.ReviewID)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, review)
}

//Testing getting review and customer
func getReviewAndCustomer(t *testing.T) GetProductReviewsRow {
	review := createTestReview(t)

	productReview, err := testQueries.GetProductReviews(context.Background(), review.ReviewID)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, productReview)

	return productReview
}

func TestProductReviews(t *testing.T) {
	productReviews := getReviewAndCustomer(t)
	require.NotEmpty(t, productReviews)
}
