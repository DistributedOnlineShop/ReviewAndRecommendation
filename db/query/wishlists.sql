-- name: CreateWishlists :one
INSERT INTO wishlists (
    wl_id,
    user_id,
    product_id
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetWishlistsList :many
SELECT 
    *
FROM 
    wishlists;

-- name: GetWishlistsByProductId :many
SELECT 
    *
FROM 
    wishlists
WHERE 
    product_id = $1;

-- name: GetWishlistsByUserId :many
SELECT 
    *
FROM 
    wishlists
WHERE 
    user_id = $1;

-- name: DeleteWishlistsByWlid :exec
DELETE FROM 
    wishlists
WHERE 
    wl_id = $1;