package main

import "testing"

func TestAverage(t *testing.T) {
	v := Average([]float64{})
	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
}
