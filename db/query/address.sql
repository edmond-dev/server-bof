-- name: CreateAddress :execresult
INSERT INTO address(
    address_customer_id,
    email,
    street,
    address_line_1,
    address_line_2,
    phone, city, state, zip_code
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
);


/* name: GetAddress :one */
SELECT * FROM address
WHERE address_customer_id = ?
ORDER BY address_customer_id
LIMIT 1;


-- name: UpdateAddress :exec
UPDATE address
SET email = ?,
    street = ?,
    address_line_1 = ?,
    address_line_2 = ?,
    phone = ?,
    city = ?,
    state = ?,
    zip_code = ?
WHERE address_customer_id = ?;

-- name: DeleteAddress :exec
DELETE FROM address
WHERE address_customer_id = ?;