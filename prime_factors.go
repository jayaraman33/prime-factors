package prime

func Factors(n int64) []int64 {
	result := []int64{}

	var i int64 = 2
	for ; n > 1; i++ {
		for ; n%i == 0; n /= i {
			result = append(result, i)
		}
	}
	return result
}
