package mocks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tamaroth/advanced-testing-techniques-in-go/cases/mocks/mocks"
)

// TestOrderService_ProcessOrderWithCustomMatcher tests the OrderService with a custom matcher for the currency argument.
func TestOrderService_ProcessOrderWithCustomMatcher(t *testing.T) {
	mockPaymentProvider := mocks.NewPaymentProvider(t)
	orderService := NewOrderService(mockPaymentProvider)

	validCurrencyMatcher := mock.MatchedBy(func(currency string) bool {
		return currency == "USD"
	})

	// Define test cases
	testCases := []struct {
		name          string
		amount        float64
		currency      string
		processResult string
		processError  error
		expectError   bool
	}{
		{
			name:          "Successful Payment in USD",
			amount:        100.0,
			currency:      "USD",
			processResult: "payment-id-123",
			processError:  nil,
			expectError:   false,
		},
		{
			name:          "Payment Failed due to Non-USD Currency",
			amount:        50.0,
			currency:      "EUR",
			processResult: "",
			processError:  fmt.Errorf("unsupported currency"),
			expectError:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockPaymentProvider.
				EXPECT().
				ProcessPayment(tc.amount, validCurrencyMatcher).
				Return(tc.processResult, tc.processError).
				Maybe()

			result, err := orderService.ProcessOrder(tc.amount, tc.currency)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.processResult, result)
			}

			mockPaymentProvider.AssertExpectations(t)
		})

		mockPaymentProvider.ExpectedCalls = nil
	}
}
