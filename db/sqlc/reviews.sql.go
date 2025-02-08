// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: reviews.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createReviews = `-- name: CreateReviews :one
INSERT INTO reviews (
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
    $5
) RETURNING review_id, user_id, product_id, rating, comment, status, created_at, updated_at
`

type CreateReviewsParams struct {
	UserID    uuid.UUID   `json:"user_id"`
	ProductID string      `json:"product_id"`
	Rating    pgtype.Int4 `json:"rating"`
	Comment   string      `json:"comment"`
	Status    string      `json:"status"`
}

func (q *Queries) CreateReviews(ctx context.Context, arg CreateReviewsParams) (Review, error) {
	row := q.db.QueryRow(ctx, createReviews,
		arg.UserID,
		arg.ProductID,
		arg.Rating,
		arg.Comment,
		arg.Status,
	)
	var i Review
	err := row.Scan(
		&i.ReviewID,
		&i.UserID,
		&i.ProductID,
		&i.Rating,
		&i.Comment,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteReview = `-- name: DeleteReview :exec
DELETE FROM 
    reviews
WHERE 
    review_id = $1
`

func (q *Queries) DeleteReview(ctx context.Context, reviewID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteReview, reviewID)
	return err
}

const getReviewByProductId = `-- name: GetReviewByProductId :one
SELECT 
    review_id, user_id, product_id, rating, comment, status, created_at, updated_at 
FROM 
    reviews
WHERE 
    product_id = $1
`

func (q *Queries) GetReviewByProductId(ctx context.Context, productID string) (Review, error) {
	row := q.db.QueryRow(ctx, getReviewByProductId, productID)
	var i Review
	err := row.Scan(
		&i.ReviewID,
		&i.UserID,
		&i.ProductID,
		&i.Rating,
		&i.Comment,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getReviewByUserId = `-- name: GetReviewByUserId :one
SELECT 
    review_id, user_id, product_id, rating, comment, status, created_at, updated_at 
FROM 
    reviews
WHERE 
    user_id = $1
`

func (q *Queries) GetReviewByUserId(ctx context.Context, userID uuid.UUID) (Review, error) {
	row := q.db.QueryRow(ctx, getReviewByUserId, userID)
	var i Review
	err := row.Scan(
		&i.ReviewID,
		&i.UserID,
		&i.ProductID,
		&i.Rating,
		&i.Comment,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getReviewsList = `-- name: GetReviewsList :many
SELECT 
    review_id, user_id, product_id, rating, comment, status, created_at, updated_at 
FROM 
    reviews
`

func (q *Queries) GetReviewsList(ctx context.Context) ([]Review, error) {
	rows, err := q.db.Query(ctx, getReviewsList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Review{}
	for rows.Next() {
		var i Review
		if err := rows.Scan(
			&i.ReviewID,
			&i.UserID,
			&i.ProductID,
			&i.Rating,
			&i.Comment,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateReviewStatus = `-- name: UpdateReviewStatus :one
UPDATE 
    reviews
SET 
    status = $2
WHERE 
    review_id = $1 RETURNING review_id, user_id, product_id, rating, comment, status, created_at, updated_at
`

type UpdateReviewStatusParams struct {
	ReviewID uuid.UUID `json:"review_id"`
	Status   string    `json:"status"`
}

func (q *Queries) UpdateReviewStatus(ctx context.Context, arg UpdateReviewStatusParams) (Review, error) {
	row := q.db.QueryRow(ctx, updateReviewStatus, arg.ReviewID, arg.Status)
	var i Review
	err := row.Scan(
		&i.ReviewID,
		&i.UserID,
		&i.ProductID,
		&i.Rating,
		&i.Comment,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
