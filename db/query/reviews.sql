-- name: CreateReviews :one
INSERT INTO reviews (
    review_id,
    USER_ID,
    PRODUCT_ID,
    RATING,
    COMMENT,
    STATUS
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: GetReviewsList :many
SELECT 
    * 
FROM 
    reviews;

-- name: GetReviewByUserId :one
SELECT 
    * 
FROM 
    reviews
WHERE 
    user_id = $1;

-- name: GetReviewByProductId :one
SELECT 
    * 
FROM 
    reviews
WHERE 
    product_id = $1;

-- name: UpdateReviewStatus :one
UPDATE 
    reviews
SET 
    status = $2,
    updated_at = NOW()
WHERE 
    review_id = $1 RETURNING *;

-- name: DeleteReview :exec
DELETE FROM 
    reviews
WHERE 
    review_id = $1;