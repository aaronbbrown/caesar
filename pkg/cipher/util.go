package cipher

func saneModInt(x, y int) int {
	result := x % y
	if result < 0 {
		result += y
	}
	return result
}

func letters() []rune {
	runes := make([]rune, 0, 26)
	for r := 'a'; r <= 'z'; r++ {
		runes = append(runes, r)
	}
	return runes
}
