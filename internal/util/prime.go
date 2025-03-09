package util

func IsPrime(value int) bool {
	for i := value - 1; i > 1; i-- {
		if value%i == 0 {
			return false
		}
	}
	return true
}
