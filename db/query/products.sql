/* name: CreateProduct :execresult */
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
);

/* name: UpdateProduct :exec */
UPDATE products
SET image_url_public_id = ?,
    image_url_secure_id = ?,
    product_name = ?,
    product_description = ?,
    price = ?,
    quantity_in_stock = ?
WHERE product_id = ?;


/* name: DeleteProduct :exec */
DELETE FROM products
WHERE product_id = ?;

/* name: GetProduct :one */
SELECT * FROM products p
WHERE p.product_id = ?
LIMIT 1;

/* name: ListProducts :many */
SELECT * FROM products
ORDER BY product_name;


/* name: GetCategoryAndProducts :many */
SELECT * FROM categories c
LEFT OUTER JOIN products p on c.category_id = p.category_id
WHERE c.category_id = ?;