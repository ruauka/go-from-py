package functions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDateDiff(t *testing.T) {
	TestTable := []struct {
		actualPaymentDate time.Time
		paymentDate       time.Time
		expected          int
		testName          string
	}{
		{
			actualPaymentDate: time.Date(2021, 01, 15, 0, 0, 0, 0, time.UTC),
			paymentDate:       time.Date(2021, 01, 15, 0, 0, 0, 0, time.UTC),
			expected:          0,
			testName:          "test-1-diff=0",
		},
		{
			actualPaymentDate: time.Date(2021, 01, 20, 0, 0, 0, 0, time.UTC),
			paymentDate:       time.Date(2021, 01, 15, 0, 0, 0, 0, time.UTC),
			expected:          5,
			testName:          "test-2-diff=5",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := DateDiff(testCase.paymentDate, testCase.actualPaymentDate)
			require.Equal(t, testCase.expected, actual)
		})
	}
}
