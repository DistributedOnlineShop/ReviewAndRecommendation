package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"ReviewAndRecommendation/util"
)

func CreateRandomWishlists(t *testing.T) Wishlist {
	data := CreateWishlistsParams{
		WlID:      util.CreateUUID(),
		UserID:    util.CreateUUID(),
		ProductID: util.GenerateProductID(),
	}

	wl, err := testStore.CreateWishlists(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, wl)
	require.Equal(t, wl.WlID, data.WlID)
	require.Equal(t, wl.UserID, data.UserID)
	require.Equal(t, wl.ProductID, data.ProductID)
	require.NotZero(t, wl.CreatedAt)

	return wl
}

func TestCreateWishlists(t *testing.T) {
	CreateRandomWishlists(t)
}

func TestDeleteWishlistsByWlid(t *testing.T) {
	wl := CreateRandomWishlists(t)

	err := testStore.DeleteWishlistsByWlid(context.Background(), wl.WlID)
	require.NoError(t, err)
}

func TestGetWishlistsByProductId(t *testing.T) {
	wl := CreateRandomWishlists(t)

	wls, err := testStore.GetWishlistsByProductId(context.Background(), wl.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, wls)
}

func TestGetWishlistsByUserId(t *testing.T) {
	wl := CreateRandomWishlists(t)

	wls, err := testStore.GetWishlistsByUserId(context.Background(), wl.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, wls)
}

func TestGetWishlistsList(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomWishlists(t)
	}

	wls, err := testStore.GetWishlistsList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, wls)
	require.GreaterOrEqual(t, len(wls), 10)
}
