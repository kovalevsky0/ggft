package opc_test

import (
	"go-sandbox/opc"
	"testing"
  "fmt"
)

func TestCalculateOrderPrice(t *testing.T) {
  type testCase struct {
    itemPrice, taxRate float64
    quantity int
    discount bool
    want string
  }

  testCases := []testCase{
    {
      itemPrice: 2.99,
      taxRate: 0.07,
      quantity: 3,
      discount: true,
      want: "The final price is: €8.64",
    },
    {
      itemPrice: 2.99,
      taxRate: 0.07,
      quantity: 3,
      discount: false,
      want: "The final price is: €9.60",
    },
  }

  for _, tc := range testCases {
    tcName := fmt.Sprintf(
      "itemPrice = %.2f, quantity = %d, taxRate = %.2f, discount = %t",
      tc.itemPrice,
      tc.quantity,
      tc.taxRate,
      tc.discount,
    )

    t.Run(tcName, func (t *testing.T)  {
      got := opc.CalculateOrderPrice(
        tc.itemPrice,
        tc.quantity,
        tc.taxRate,
        tc.discount,
      )

      if tc.want != got {
        t.Errorf("Failed! Wanted %q, got %q", tc.want, got)
      }
    })
  }
}
