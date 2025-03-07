// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateCoupons(ctx context.Context, arg CreateCouponsParams) (Coupon, error)
	CreateReviews(ctx context.Context, arg CreateReviewsParams) (Review, error)
	CreateWishlists(ctx context.Context, arg CreateWishlistsParams) (Wishlist, error)
	DeleteCoupon(ctx context.Context, couponID uuid.UUID) error
	DeleteReview(ctx context.Context, reviewID uuid.UUID) error
	DeleteWishlistsByWlid(ctx context.Context, wlID uuid.UUID) error
	GetCouponByCouponId(ctx context.Context, couponID uuid.UUID) (Coupon, error)
	GetReviewByProductId(ctx context.Context, productID string) (Review, error)
	GetReviewByUserId(ctx context.Context, userID uuid.UUID) (Review, error)
	GetReviewsList(ctx context.Context) ([]Review, error)
	GetValidCoupons(ctx context.Context) ([]Coupon, error)
	GetWishlistsByProductId(ctx context.Context, productID string) ([]Wishlist, error)
	GetWishlistsByUserId(ctx context.Context, userID uuid.UUID) ([]Wishlist, error)
	GetWishlistsList(ctx context.Context) ([]Wishlist, error)
	UpdateCouponExpiresAt(ctx context.Context, arg UpdateCouponExpiresAtParams) (Coupon, error)
	UpdateReviewStatus(ctx context.Context, arg UpdateReviewStatusParams) (Review, error)
}

var _ Querier = (*Queries)(nil)
