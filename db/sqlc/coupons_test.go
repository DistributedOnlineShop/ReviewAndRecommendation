package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"

	"ReviewAndRecommendation/util"
)

func CreateRandomCoupons(t *testing.T) Coupon {
	startAt := pgtype.Timestamp{
		Time:  util.GenerateDate().Time.Add(-100 * time.Hour),
		Valid: true,
	}
	data := CreateCouponsParams{
		CouponID:    util.CreateUUID(),
		Code:        util.GenerateText(),
		Discount:    util.GenerateNumeric(),
		MinPurchase: util.GenerateNumeric(),
		StartAt:     startAt,
		ExpiresAt:   util.GenerateDate(),
	}

	coupon, err := testStore.CreateCoupons(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, coupon.CouponID)
	require.Equal(t, data.Code, coupon.Code)
	require.Equal(t, data.Discount, coupon.Discount)
	require.Equal(t, data.MinPurchase, coupon.MinPurchase)
	require.NotZero(t, coupon.CreatedAt)

	return coupon
}

func TestCreateCoupons(t *testing.T) {
	CreateRandomCoupons(t)
}

func TestDeleteCoupon(t *testing.T) {
	coupon := CreateRandomCoupons(t)

	err := testStore.DeleteCoupon(context.Background(), coupon.CouponID)
	require.NoError(t, err)
}

func TestGetCouponByCouponId(t *testing.T) {
	coupon := CreateRandomCoupons(t)

	c, err := testStore.GetCouponByCouponId(context.Background(), coupon.CouponID)
	require.NoError(t, err)
	require.NotEmpty(t, c)
}

func TestGetValidCoupons(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomCoupons(t)
	}

	coupons, err := testStore.GetValidCoupons(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, coupons)
	require.GreaterOrEqual(t, len(coupons), 1)
}

func TestUpdateCouponExpiresAt(t *testing.T) {
	coupon := CreateRandomCoupons(t)

	newData := UpdateCouponExpiresAtParams{
		CouponID:  coupon.CouponID,
		ExpiresAt: util.GenerateDate(),
	}
	updatedCoupon, err := testStore.UpdateCouponExpiresAt(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, updatedCoupon)
	require.NotEqual(t, coupon.ExpiresAt, updatedCoupon.ExpiresAt)
	require.NotZero(t, updatedCoupon.UpdatedAt)
}
