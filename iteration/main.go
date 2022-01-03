package iteration

func Repeat(char string, numOfRepeat int) string {
	var repeated string
	for i := 0; i < numOfRepeat; i++ {
		repeated += char
	}
	return repeated
}
