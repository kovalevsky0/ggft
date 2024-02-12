package calculator_test

import (
	"go-sandbox/calculator"
	"testing"
  "fmt"
)

func TestSum(t *testing.T) {
	type testCase struct {
		a    int
		b    int
		c    int
		want int
	}

	testCases := []testCase{
		{
			a:    1,
			b:    2,
			c:    3,
			want: 6,
		},
		{
			a:    4,
			b:    5,
			c:    6,
			want: 15,
		},
		{
			a:    7,
			b:    8,
			c:    9,
			want: 24,
		},
	}

	for _, tc := range testCases {
    t.Run(fmt.Sprintf("a = %d, b = %d, c = %d", tc.a, tc.b, tc.c), func(t *testing.T)  {
      got := calculator.Sum(tc.a, tc.b, tc.c)
      if got != tc.want {
        t.Errorf("Expected %d but got %d", tc.want, got)
      }
    })
	}
}
