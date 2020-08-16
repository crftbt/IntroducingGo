package packagetesting

// finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// finds the maximum number of a series of numbers
func Maximum(xs []float64) float64 {
	var max float64
	for i, x := range xs {
		if i == 0 || x > max {
			max = x
		}
	}
	return max
}

// finds the minimum number of a series of numbers
func Minimum(xs []float64) float64 {
	var min float64
	for i, x := range xs {
		if i == 0 || x < min {
			min = x
		}
	}
	return min
}
