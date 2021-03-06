// Code generated by sqlc. DO NOT EDIT.
// source: orderDetails.sql

package db

import (
	"context"
	"database/sql"
)

const createOrderDetails = `-- name: CreateOrderDetails :execresult
INSERT INTO orderDetails(
 order_number, shipped_date, status, comments
) VALUES (
     ?, ?, ?, ?
)
`

type CreateOrderDetailsParams struct {
	OrderNumber string `json:"orderNumber"`
	ShippedDate string `json:"shippedDate"`
	Status      string `json:"status"`
	Comments    string `json:"comments"`
}

func (q *Queries) CreateOrderDetails(ctx context.Context, arg CreateOrderDetailsParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createOrderDetails,
		arg.OrderNumber,
		arg.ShippedDate,
		arg.Status,
		arg.Comments,
	)
}

const getOrderDetails = `-- name: GetOrderDetails :one
SELECT order_number, shipped_date, status, comments FROM orderDetails
WHERE order_number  = ?
`

func (q *Queries) GetOrderDetails(ctx context.Context, orderNumber string) (OrderDetail, error) {
	row := q.db.QueryRowContext(ctx, getOrderDetails, orderNumber)
	var i OrderDetail
	err := row.Scan(
		&i.OrderNumber,
		&i.ShippedDate,
		&i.Status,
		&i.Comments,
	)
	return i, err
}
