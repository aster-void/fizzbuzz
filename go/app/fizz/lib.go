package fizz

func Fizz(i int) string {
	if isFizz(i) {
		return "Fizz"
	}
	return ""
}

func isFizz(i int) bool {
	return i%3 == 0
}
