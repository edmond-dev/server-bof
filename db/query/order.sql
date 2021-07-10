-- name: CreateOrder :execresult
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
 );

-- name: GetLatestOrderNumber :one
SELECT @order_number := MAX(order_number)+1
FROM orders;

/* name: GetOrder :one */
SELECT * FROM orders
WHERE order_number = ?
ORDER BY order_number;


-- name: UpdateOrder :exec
UPDATE orders
SET customer_comments = ?
WHERE order_number = ?;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_number = ?;