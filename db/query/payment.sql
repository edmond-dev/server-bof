-- name: CreateOrderPayment :execresult
INSERT INTO payment (
    customer_id, payment_date, payment_method_id, amount
) VALUES (
     ?, ?, ?, ?
);

-- name: GetPayment :one
SELECT * FROM payment
WHERE customer_id = ?
ORDER BY  customer_id;