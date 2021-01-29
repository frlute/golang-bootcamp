package mymath

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := []struct {
		data     float64
		accuracy float64
		expect   float64
	}{
		{
			4.0,
			0,
			2.0,
		},
		{
			3,
			0,
			1.7320508075688772,
		},
	}

	for _, c := range cases {
		got, err := Sqrt(c.data, c.accuracy)
		if err != nil {
			t.Errorf("Sqrt called occur error: %v", err)
		}
		if got != c.expect {
			t.Errorf("Sqrt(%v) execute expect value: %v, actual value: %v", c.data, c.expect, got)
		}
	}

	sqrt2Cases := []struct {
		data   int
		expect int
	}{
		{
			4.0,
			2.0,
		},
		{
			3,
			1,
		},
	}

	for _, c := range sqrt2Cases {
		got := Sqrt2(c.data)
		if got != c.expect {
			t.Errorf("Sqrt(%v) execute expect value: %v, actual value: %v", c.data, c.expect, got)
		}
	}
}
