package utils

func FloatEqualsWithPrecision(a float64, b float64, eps float64) bool {
	if (a-b) < eps && (b-a) < eps {
		return true
	} else {
		return false
	}
}
