-- name: CreateReview :execresult
INSERT INTO reviews (
   review_id, product_review_id, customer_review_id, review
) VALUES (
     ?, ?, ?, ?
 );

/* name: GetReview :one */
SELECT * FROM reviews
WHERE review_id = ?
ORDER BY review_id;

/* name: GetProductReviews :one */
SELECT r.*, c.* FROM reviews r
LEFT JOIN products p on p.product_id = r.product_review_id
LEFT JOIN customers c on c.customer_id = r.customer_review_id
WHERE review_id = ?
LIMIT 1;


-- name: DeleteReview :exec
DELETE FROM reviews
WHERE review_id = ?;

-- name: UpdateReview :exec
UPDATE reviews
SET review = ?
WHERE review_id = ?;
