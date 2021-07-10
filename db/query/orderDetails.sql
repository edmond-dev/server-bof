-- name: CreateOrderDetails :execresult
INSERT INTO orderDetails(
 order_number, shipped_date, status, comments
) VALUES (
     ?, ?, ?, ?
);

-- name: GetOrderDetails :one
SELECT * FROM orderDetails
WHERE order_number  = ?;