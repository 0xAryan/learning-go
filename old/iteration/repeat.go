package iteration


const repeatCount = 5

func Repeat(character string, number int) string{
	var repeated string
	for i := 0; i < number; i++ {
		// repeated = repeated + character
		repeated += character
	}
	return repeated
}