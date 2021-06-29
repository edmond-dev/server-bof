/* name: CreateUser :execresult */
INSERT INTO users (
    user_id,
    first_name,
    last_name,
    email,
    password

) VALUES (
        ?, ?, ?, ?, ?
 );

/* name: GetUser :one */
SELECT * FROM users
WHERE id = ? LIMIT 1;

/* name: ListUser :many */
SELECT * FROM users
ORDER BY first_name;

/* name: DeleteUser :exec */
DELETE FROM users
WHERE id = ?;


-- name: UpdateUser :exec
UPDATE users
SET first_name = ?, last_name = ?, email = ?, password = ?
WHERE id = ?;