/* name: CreateProduct :execresult */
INSERT INTO products (
  product_id,
  image_url_public_id,
  image_url_secure_id,
  product_name,
  product_description,
  price

) VALUES (
    ?, ?, ?, ?, ?, ?
);

/* name: GetProduct :one */
SELECT * FROM products
WHERE id = ?;


/* name: ListProducts :many */
SELECT * FROM products
ORDER BY product_name;

/* name: DeleteProduct :exec */
DELETE FROM products
WHERE id = ?;


-- name: UpdateProduct :exec
UPDATE products
SET image_url_public_id = ?,
    image_url_secure_id = ?,
    product_name = ?,
    product_description = ?,
    price = ?

WHERE id = ?;