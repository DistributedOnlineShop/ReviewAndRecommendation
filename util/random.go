package util

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateText() pgtype.Text {
	return pgtype.Text{
		String: gofakeit.Word(),
		Valid:  true,
	}
}

func GenerateDate() pgtype.Timestamp {
	daysOffset := rand.IntN(365) - 180
	return pgtype.Timestamp{
		Time:  time.Now().Add(time.Duration(daysOffset) * 24 * time.Hour),
		Valid: true,
	}
}

func GenerateNumeric() pgtype.Numeric {
	intPart := rand.IntN(100000)
	fracPart := rand.IntN(100)
	value := int64(intPart) + int64(fracPart)/100.0

	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   -2,
		Valid: true,
	}
}

func GenerateStatus() string {
	statuses := []string{
		"ACTIVE",
		"INACTIVE",
		"PENDING",
		"SUSPENDED",
		"BLOCKED",
		"ARCHIVED",
	}
	return statuses[rand.IntN(len(statuses))]
}

func GenerateInt32() int32 {
	return rand.Int32N(1000) + 1
}

func GenerateBool() pgtype.Bool {
	return pgtype.Bool{
		Bool:  gofakeit.Bool(),
		Valid: true,
	}
}

func GenerateProductID() string {
	number := rand.IntN(999999) + 1
	return fmt.Sprintf("P%06d", number)
}

func GenerateReviewStatus() string {
	commentStatuses := []string{
		"PENDING",
		"APPROVED",
		"REJECTED",
		"REPORTED",
		"DELETED",
	}
	return commentStatuses[rand.IntN(len(commentStatuses))]
}
