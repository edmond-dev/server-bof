-- name: CreateCustomer :execresult
INSERT INTO customers (
  customer_id, first_name, last_name, email, role, password
) VALUES (
       ?, ?, ?, ?, ?, ?
 );

/* name: UpdateCustomer :exec */
UPDATE customers
SET first_name = ?, last_name = ?, email = ?

WHERE customer_id = ?;


-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE customer_id = ?;


-- name: GetCustomerWithId :one
SELECT * FROM customers
WHERE customer_id = ? LIMIT 1;

-- name: GetCustomerWithEmail :one
SELECT * FROM customers
WHERE email = ? LIMIT 1;

-- name: GetCustomerOrderDetailsAndAddr :one
SELECT order_number, c.customer_id, product_id, order_category_name,
       quantity_ordered, price_each, ordered_date,
       customer_comments,
       created_at, address_customer_id, street,
       address_line_1, address_line_2, phone, city, state, zip_code

FROM orders o
join customers c on c.customer_id = o.customer_id
join address a on c.customer_id = a.address_customer_id
WHERE c.customer_id = ?;
