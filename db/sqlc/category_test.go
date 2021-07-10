package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"log"
	"server-bof/util"
	"strings"
	"testing"
)

//Testing creating a test category
func createCategoryTest(t *testing.T) Category {

	//Category params
	arg := CreateCategoryParams{
		CategoryID:              strings.ToUpper(uuid.New().String()),
		CategoryName:            util.RandomString(7) + "Flower - Category Name",
		CategoryHtmlDescription: "some category description",
		Image:                   "some image category image",
	}

	res, err := testQueries.CreateCategory(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	categoryId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	category, err := testQueries.GetCategory(context.Background(), arg.CategoryID)

	log.Println(categoryId)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	return category
}

//Testing creating a category
func TestCreateCategory(t *testing.T) {
	createCategoryTest(t)
}

//Testing getting a category
func TestGetCategory(t *testing.T) {
	category := createCategoryTest(t)

	category2, err := testQueries.GetCategory(context.Background(), category.CategoryID)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category.CategoryID, category2.CategoryID)
	require.Equal(t, category.CategoryName, category2.CategoryName)
	require.Equal(t, category.CategoryHtmlDescription, category2.CategoryHtmlDescription)

}

//Testing getting a categories
func TestGetCategories(t *testing.T) {
	categories, err := testQueries.GetCategories(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, categories)
}

//Testing deleting a category
func TestDeleteCategory(t *testing.T) {
	category := createCategoryTest(t)

	err := testQueries.DeleteCategory(context.Background(), category.CategoryID)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, category)
}

//Testing updating a category
func TestUpdateCategory(t *testing.T) {
	category := createCategoryTest(t)

	arg := UpdateCategoryParams{
		CategoryName:            category.CategoryName,
		CategoryHtmlDescription: category.CategoryHtmlDescription,
		Image:                   category.Image,
		CategoryID:              category.CategoryID,
	}
	err := testQueries.UpdateCategory(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, category)
}
