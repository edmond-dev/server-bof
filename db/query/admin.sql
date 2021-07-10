-- name: GetProductStatsPerCategory :one
SELECT c.category_name,
       COUNT(*) AS TotalProductsPerCategory
FROM categories c
INNER JOIN products p on c.category_id = p.category_id
WHERE c.category_id  = ?
GROUP BY c.category_name;

-- name: ProductsAvailable :one
SELECT product_name, COUNT(*) AS TotalProductsAvailable FROM products
GROUP BY product_name;

-- name: GetTotalRevenue :one
SELECT p.amount, SUM(p.amount) AS totalRevenue
FROM payment p
GROUP BY amount;

-- name: GetTotalCustomers :one
SELECT customer_id, COUNT(*) AS TotalNumberOfCustomers
FROM customers
GROUP BY customer_id;

-- name: TotalAmountOfOrdersByACustomer :one
SELECT customer_id, sum(quantity_ordered * orders.price_each) AS CutomerTotal
FROM orders
INNER JOIN orderDetails oD on orders.order_number = oD.order_number
GROUP BY customer_id;