// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: coupons.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createCoupons = `-- name: CreateCoupons :one
INSERT INTO coupons (
    code,
    discount,
    min_purchase,
    expires_at
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING coupon_id, code, discount, min_purchase, expires_at, created_at, updated_at
`

type CreateCouponsParams struct {
	Code        pgtype.Text      `json:"code"`
	Discount    pgtype.Numeric   `json:"discount"`
	MinPurchase pgtype.Numeric   `json:"min_purchase"`
	ExpiresAt   pgtype.TIMESTAMP(0) `json:"expires_at"`
}

func (q *Queries) CreateCoupons(ctx context.Context, arg CreateCouponsParams) (Coupon, error) {
	row := q.db.QueryRow(ctx, createCoupons,
		arg.Code,
		arg.Discount,
		arg.MinPurchase,
		arg.ExpiresAt,
	)
	var i Coupon
	err := row.Scan(
		&i.CouponID,
		&i.Code,
		&i.Discount,
		&i.MinPurchase,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCoupon = `-- name: DeleteCoupon :exec
DELETE FROM 
    coupons 
WHERE 
    coupon_id = $1 RETURNING coupon_id, code, discount, min_purchase, expires_at, created_at, updated_at
`

func (q *Queries) DeleteCoupon(ctx context.Context, couponID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteCoupon, couponID)
	return err
}

const getCouponByCouponId = `-- name: GetCouponByCouponId :one
SELECT 
    coupon_id, code, discount, min_purchase, expires_at, created_at, updated_at 
FROM 
    coupons 
WHERE 
    coupon_id = $1
`

func (q *Queries) GetCouponByCouponId(ctx context.Context, couponID uuid.UUID) (Coupon, error) {
	row := q.db.QueryRow(ctx, getCouponByCouponId, couponID)
	var i Coupon
	err := row.Scan(
		&i.CouponID,
		&i.Code,
		&i.Discount,
		&i.MinPurchase,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getValidCoupons = `-- name: GetValidCoupons :many
SELECT 
    coupon_id, code, discount, min_purchase, expires_at, created_at, updated_at
FROM 
    coupons
WHERE 
    expires_at >= NOW()
`

func (q *Queries) GetValidCoupons(ctx context.Context) ([]Coupon, error) {
	rows, err := q.db.Query(ctx, getValidCoupons)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Coupon{}
	for rows.Next() {
		var i Coupon
		if err := rows.Scan(
			&i.CouponID,
			&i.Code,
			&i.Discount,
			&i.MinPurchase,
			&i.ExpiresAt,
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

const updateExpiresAt = `-- name: UpdateExpiresAt :one
UPDATE 
    coupons 
SET 
    expires_at = $1
WHERE 
    coupon_id = $2 RETURNING coupon_id, code, discount, min_purchase, expires_at, created_at, updated_at
`

type UpdateExpiresAtParams struct {
	ExpiresAt pgtype.TIMESTAMP(0) `json:"expires_at"`
	CouponID  uuid.UUID        `json:"coupon_id"`
}

func (q *Queries) UpdateExpiresAt(ctx context.Context, arg UpdateExpiresAtParams) (Coupon, error) {
	row := q.db.QueryRow(ctx, updateExpiresAt, arg.ExpiresAt, arg.CouponID)
	var i Coupon
	err := row.Scan(
		&i.CouponID,
		&i.Code,
		&i.Discount,
		&i.MinPurchase,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
