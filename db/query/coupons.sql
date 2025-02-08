-- name: CreateCoupons :one
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
    expires_at >= NOW();

-- name: UpdateExpiresAt :one
UPDATE 
    coupons 
SET 
    expires_at = $1
WHERE 
    coupon_id = $2 RETURNING *;

-- name: DeleteCoupon :exec
DELETE FROM 
    coupons 
WHERE 
    coupon_id = $1 RETURNING *;