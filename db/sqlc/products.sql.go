// Code generated by sqlc. DO NOT EDIT.
// source: products.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createProduct = `-- name: CreateProduct :execresult
INSERT INTO products (
  product_id,
  category_id,
  image_url_public_id,
  image_url_secure_id,
  product_name,
  product_description,
  price,
  quantity_in_stock

) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateProductParams struct {
	ProductID          string `json:"productID"`
	CategoryID         string `json:"categoryID"`
	ImageUrlPublicID   string `json:"imageUrlPublicID"`
	ImageUrlSecureID   string `json:"imageUrlSecureID"`
	ProductName        string `json:"productName"`
	ProductDescription string `json:"productDescription"`
	Price              int64  `json:"price"`
	QuantityInStock    string `json:"quantityInStock"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProduct,
		arg.ProductID,
		arg.CategoryID,
		arg.ImageUrlPublicID,
		arg.ImageUrlSecureID,
		arg.ProductName,
		arg.ProductDescription,
		arg.Price,
		arg.QuantityInStock,
	)
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = ?
`

func (q *Queries) DeleteProduct(ctx context.Context, productID string) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, productID)
	return err
}

const getCategoryAndProducts = `-- name: GetCategoryAndProducts :many
SELECT c.category_id, category_name, category_html_description, image, product_id, p.category_id, image_url_public_id, image_url_secure_id, product_name, product_description, price, quantity_in_stock, created_at FROM categories c
LEFT OUTER JOIN products p on c.category_id = p.category_id
WHERE c.category_id = ?
`

type GetCategoryAndProductsRow struct {
	CategoryID              string    `json:"categoryID"`
	CategoryName            string    `json:"categoryName"`
	CategoryHtmlDescription string    `json:"categoryHtmlDescription"`
	Image                   string    `json:"image"`
	ProductID               string    `json:"productID"`
	CategoryID_2            string    `json:"categoryID2"`
	ImageUrlPublicID        string    `json:"imageUrlPublicID"`
	ImageUrlSecureID        string    `json:"imageUrlSecureID"`
	ProductName             string    `json:"productName"`
	ProductDescription      string    `json:"productDescription"`
	Price                   int64     `json:"price"`
	QuantityInStock         string    `json:"quantityInStock"`
	CreatedAt               time.Time `json:"createdAt"`
}

func (q *Queries) GetCategoryAndProducts(ctx context.Context, categoryID string) ([]GetCategoryAndProductsRow, error) {
	rows, err := q.db.QueryContext(ctx, getCategoryAndProducts, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCategoryAndProductsRow
	for rows.Next() {
		var i GetCategoryAndProductsRow
		if err := rows.Scan(
			&i.CategoryID,
			&i.CategoryName,
			&i.CategoryHtmlDescription,
			&i.Image,
			&i.ProductID,
			&i.CategoryID_2,
			&i.ImageUrlPublicID,
			&i.ImageUrlSecureID,
			&i.ProductName,
			&i.ProductDescription,
			&i.Price,
			&i.QuantityInStock,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProduct = `-- name: GetProduct :one
SELECT product_id, category_id, image_url_public_id, image_url_secure_id, product_name, product_description, price, quantity_in_stock, created_at FROM products p
WHERE p.product_id = ?
LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, productID string) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.CategoryID,
		&i.ImageUrlPublicID,
		&i.ImageUrlSecureID,
		&i.ProductName,
		&i.ProductDescription,
		&i.Price,
		&i.QuantityInStock,
		&i.CreatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT product_id, category_id, image_url_public_id, image_url_secure_id, product_name, product_description, price, quantity_in_stock, created_at FROM products
ORDER BY product_name
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.CategoryID,
			&i.ImageUrlPublicID,
			&i.ImageUrlSecureID,
			&i.ProductName,
			&i.ProductDescription,
			&i.Price,
			&i.QuantityInStock,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET image_url_public_id = ?,
    image_url_secure_id = ?,
    product_name = ?,
    product_description = ?,
    price = ?,
    quantity_in_stock = ?
WHERE product_id = ?
`

type UpdateProductParams struct {
	ImageUrlPublicID   string `json:"imageUrlPublicID"`
	ImageUrlSecureID   string `json:"imageUrlSecureID"`
	ProductName        string `json:"productName"`
	ProductDescription string `json:"productDescription"`
	Price              int64  `json:"price"`
	QuantityInStock    string `json:"quantityInStock"`
	ProductID          string `json:"productID"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct,
		arg.ImageUrlPublicID,
		arg.ImageUrlSecureID,
		arg.ProductName,
		arg.ProductDescription,
		arg.Price,
		arg.QuantityInStock,
		arg.ProductID,
	)
	return err
}