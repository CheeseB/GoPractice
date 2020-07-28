package iteration

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 7; i++ {
		repeated += character
	}
	return repeated
}
