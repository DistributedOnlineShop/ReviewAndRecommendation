package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"ReviewAndRecommendation/util"
)

func CreateRandomReviews(t *testing.T) Review {
	data := CreateReviewsParams{
		ReviewID:  util.CreateUUID(),
		UserID:    util.CreateUUID(),
		ProductID: util.GenerateProductID(),
		Rating:    util.GenerateNumeric(),
		Comment:   gofakeit.Word(),
		Status:    util.GenerateReviewStatus(),
	}

	review, err := testStore.CreateReviews(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, review)
	require.Equal(t, data.ReviewID, review.ReviewID)
	require.Equal(t, data.UserID, review.UserID)
	require.Equal(t, data.ProductID, review.ProductID)
	require.Equal(t, data.Rating, review.Rating)
	require.Equal(t, data.Comment, review.Comment)
	require.NotEmpty(t, review.CreatedAt)

	return review

}

func TestCreateReviews(t *testing.T) {
	CreateRandomReviews(t)
}

func TestDeleteReview(t *testing.T) {
	review := CreateRandomReviews(t)

	err := testStore.DeleteReview(context.Background(), review.ReviewID)
	require.NoError(t, err)
}

func TestGetReviewByProductId(t *testing.T) {
	review := CreateRandomReviews(t)

	got, err := testStore.GetReviewByProductId(context.Background(), review.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, got)
}

func TestGetReviewByUserId(t *testing.T) {
	review := CreateRandomReviews(t)

	got, err := testStore.GetReviewByUserId(context.Background(), review.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, got)
}

func TestGetReviewsList(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomReviews(t)
	}

	list, err := testStore.GetReviewsList(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(list), 10)
}

func TestUpdateReviewStatus(t *testing.T) {
	review := CreateRandomReviews(t)

	newData := UpdateReviewStatusParams{
		ReviewID: review.ReviewID,
		Status:   util.GenerateReviewStatus(),
	}

	got, err := testStore.UpdateReviewStatus(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.Equal(t, review.ReviewID, got.ReviewID)
	require.NotZero(t, got.UpdatedAt)
}
