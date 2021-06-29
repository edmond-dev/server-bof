/* name: CreateReview :execresult */
INSERT INTO reviews (
   review_id,
   user_review_id,
   review
) VALUES (
     ?, ?, ?
 );

/* name: GetReview :one */
SELECT * FROM reviews
WHERE review_id = ?;


/* name: DeleteReview :exec */
DELETE FROM reviews
WHERE id = ?;


-- name: UpdateReview :exec
UPDATE reviews
SET review = ?
WHERE id = ?;
