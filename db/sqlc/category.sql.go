// Code generated by sqlc. DO NOT EDIT.
// source: category.sql

package db

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :execresult
INSERT INTO categories (
    category_id, category_name, category_html_description, image
)
VALUE (
    ?, ?, ?, ?
)
`

type CreateCategoryParams struct {
	CategoryID              string `json:"categoryID"`
	CategoryName            string `json:"categoryName"`
	CategoryHtmlDescription string `json:"categoryHtmlDescription"`
	Image                   string `json:"image"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCategory,
		arg.CategoryID,
		arg.CategoryName,
		arg.CategoryHtmlDescription,
		arg.Image,
	)
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
WHERE category_id = ?
`

func (q *Queries) DeleteCategory(ctx context.Context, categoryID string) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, categoryID)
	return err
}

const getCategories = `-- name: GetCategories :many
SELECT category_id, category_name, category_html_description, image FROM categories
ORDER BY category_id
`

func (q *Queries) GetCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.CategoryID,
			&i.CategoryName,
			&i.CategoryHtmlDescription,
			&i.Image,
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

const getCategory = `-- name: GetCategory :one
SELECT category_id, category_name, category_html_description, image FROM categories
WHERE category_id = ?
`

func (q *Queries) GetCategory(ctx context.Context, categoryID string) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, categoryID)
	var i Category
	err := row.Scan(
		&i.CategoryID,
		&i.CategoryName,
		&i.CategoryHtmlDescription,
		&i.Image,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories
SET category_name = ?,
    category_html_description = ?,
    image = ?
WHERE category_id = ?
`

type UpdateCategoryParams struct {
	CategoryName            string `json:"categoryName"`
	CategoryHtmlDescription string `json:"categoryHtmlDescription"`
	Image                   string `json:"image"`
	CategoryID              string `json:"categoryID"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory,
		arg.CategoryName,
		arg.CategoryHtmlDescription,
		arg.Image,
		arg.CategoryID,
	)
	return err
}