package iteration

const repeatCount = 5

// Repeat returns character repeated 5 times.
func Repeat(char string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}

// CharRepeat returns character repeated inputed number of time.
func CharRepeat(char string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += char
	}
	return repeated
}

