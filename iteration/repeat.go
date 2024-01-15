package iteration


// Pass the character and no of time to repeat it and it will return that string
func Repeat(character string, no int) string {
	var repeated string

	for i:= 0; i < no; i++ {
		repeated = repeated + character
	}

	return repeated
}