package utils

// Smart batch/goroutine calculator
func DecideBatchSizeAndConcurrency(n int) (int, int) {
	switch {
	case n <= 10_000:
		return 500, 2
	case n <= 100_000:
		return 1000, 5
	case n <= 1_000_000:
		return 2000, 10
	default:
		return 5000, 20
	}
}
