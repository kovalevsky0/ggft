package opc

import "fmt"

const discountRate float64 = 0.9

func CalculateOrderPrice(
  itemPrice float64,
  quantity int,
  taxRate float64,
  discount bool,
) string {
  totalPrice := itemPrice * float64(quantity)

  if discount {
    totalPrice = totalPrice * discountRate
  }

  taxAmount := totalPrice * taxRate
  finalPrice := totalPrice + taxAmount

  return fmt.Sprintf("The final price is: €%.2f", finalPrice)
}
