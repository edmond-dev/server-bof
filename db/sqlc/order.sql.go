// Code generated by sqlc. DO NOT EDIT.
// source: order.sql

package db

import (
	"context"
	"database/sql"
)

const createOrder = `-- name: CreateOrder :execresult
INSERT INTO orders (
    order_number,
    customer_id,
    product_id,
    order_category_name,
    quantity_ordered,
    price_each,
    customer_comments

) VALUES (
   ?, ?, ?, ?, ?, ?, ?
 )
`

type CreateOrderParams struct {
	OrderNumber       string `json:"orderNumber"`
	CustomerID        string `json:"customerID"`
	ProductID         string `json:"productID"`
	OrderCategoryName string `json:"orderCategoryName"`
	QuantityOrdered   int32  `json:"quantityOrdered"`
	PriceEach         int32  `json:"priceEach"`
	CustomerComments  string `json:"customerComments"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createOrder,
		arg.OrderNumber,
		arg.CustomerID,
		arg.ProductID,
		arg.OrderCategoryName,
		arg.QuantityOrdered,
		arg.PriceEach,
		arg.CustomerComments,
	)
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_number = ?
`

func (q *Queries) DeleteOrder(ctx context.Context, orderNumber string) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, orderNumber)
	return err
}

const getLatestOrderNumber = `-- name: GetLatestOrderNumber :one
SELECT @order_number := MAX(order_number)+1
FROM orders
`

func (q *Queries) GetLatestOrderNumber(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getLatestOrderNumber)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getOrder = `-- name: GetOrder :one
SELECT order_number, customer_id, product_id, order_category_name, quantity_ordered, price_each, customer_comments, ordered_date FROM orders
WHERE order_number = ?
ORDER BY order_number
`

func (q *Queries) GetOrder(ctx context.Context, orderNumber string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, orderNumber)
	var i Order
	err := row.Scan(
		&i.OrderNumber,
		&i.CustomerID,
		&i.ProductID,
		&i.OrderCategoryName,
		&i.QuantityOrdered,
		&i.PriceEach,
		&i.CustomerComments,
		&i.OrderedDate,
	)
	return i, err
}

const updateOrder = `-- name: UpdateOrder :exec
UPDATE orders
SET customer_comments = ?
WHERE order_number = ?
`

type UpdateOrderParams struct {
	CustomerComments string `json:"customerComments"`
	OrderNumber      string `json:"orderNumber"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) error {
	_, err := q.db.ExecContext(ctx, updateOrder, arg.CustomerComments, arg.OrderNumber)
	return err
}