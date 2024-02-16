package iteration

const defaultRepeatCount = 5

// Repeat takes a string and a count and returns a string with the input string repeated for count times
func Repeat(char string, count int) string {
	var repeated string

	if count == 0 {
		count = defaultRepeatCount
	}

	for i := 0; i < count; i++ {
		repeated += char
	}
	return repeated
}
