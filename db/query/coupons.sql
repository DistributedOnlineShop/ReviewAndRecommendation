-- name: CreateCoupons :one
INSERT INTO coupons (
    coupon_id,
    code,
    discount,
    min_purchase,
    start_at,
    expires_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: GetCouponByCouponId :one
SELECT 
    * 
FROM 
    coupons 
WHERE 
    coupon_id = $1;

-- name: GetValidCoupons :many
SELECT 
    *
FROM 
    coupons
WHERE 
    start_at < NOW() AND expires_at >= NOW();

-- name: UpdateCouponExpiresAt :one
UPDATE 
    coupons 
SET 
    expires_at = $2,
    updated_at = NOW()
WHERE 
    coupon_id = $1 RETURNING *;

-- name: DeleteCoupon :exec
DELETE FROM 
    coupons 
WHERE 
    coupon_id = $1 RETURNING *;