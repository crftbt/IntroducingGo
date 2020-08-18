//chapter 9 exercise 2 - Write a series of tests for hte Min and Max functions from chapter 8.

package packagetesting

import "testing"

func TestPackagetesting(t *testing.T) {
	cases := []struct {
		xs       []float64
		max, min float64
	}{
		{
			xs:  []float64{3, 5, 2, 1, 7, 9},
			max: 9,
			min: 1,
		},
		{
			xs:  []float64{},
			max: 0,
			min: 0,
		},
	}

	for _, c := range cases {
		max := Maximum(c.xs)
		if max != c.max {
			t.Errorf("expected %f got %f", c.max, max)
		}
		min := Minimum(c.xs)
		if min != c.min {
			t.Errorf("expected %f got %f", c.min, min)
		}
	}
}
