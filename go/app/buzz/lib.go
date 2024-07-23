package buzz

func Buzz(i int) string {
	if isBuzz(i) {
		return "Buzz"
	}
	return ""
}

func isBuzz(i int) bool {
	return i%5 == 0
}
