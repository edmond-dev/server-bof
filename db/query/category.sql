-- name: CreateCategory :execresult
INSERT INTO categories (
    category_id, category_name, category_html_description, image
)
VALUE (
    ?, ?, ?, ?
);

-- name: UpdateCategory :exec
UPDATE categories
SET category_name = ?,
    category_html_description = ?,
    image = ?
WHERE category_id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE category_id = ?;

-- name: GetCategories :many
SELECT * FROM categories
ORDER BY category_id;

-- name: GetCategory :one
SELECT * FROM categories
WHERE category_id = ?;
