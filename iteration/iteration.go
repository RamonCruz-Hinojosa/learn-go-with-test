package iteration

func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
		//using := is declaring and assigning in one go while = is just assigning := would not work on this line
	}
	return repeated
}
