package request

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidateStruct(t *testing.T) {
	req := NewTestReq()
	t.Run("validation error", func(t *testing.T) {
		err := ValidateStruct(req)
		fmt.Println(err)
		require.EqualError(t, err, "Key: 'Request.Loan[0].Payment' Error:Field validation for 'Payment' failed on the 'required' tag")
	})
	t.Run("success validation", func(t *testing.T) {
		req.Loan[0].Payment = 10
		err := ValidateStruct(req)
		require.Equal(t, nil, err)
	})
}

func NewTestReq() *Request {
	paymentDateBalance := 10

	return &Request{
		Name:    "Ivan",
		Surname: "Ivanov",
		Account: []Account{
			{
				PaymentDateBalance: &paymentDateBalance,
			},
		},
		Loan: []Loan{
			{

				PaymentDate:       time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
				ActualPaymentDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			},
		},
	}
}
