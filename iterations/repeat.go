package iterations

//Repeat a number of time a character
func Repeat(character string, number int) (repeated string) {
	for i := 0; i < number; i++ {
		repeated += character
	}

	return
}
