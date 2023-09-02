package iteration

func Repeat(text string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated += text
	}

	return repeated
}
