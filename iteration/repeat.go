package iteration

// Repeats given string rc times
func Repeat(c string, rc int) string {
	var r string

	for i := 0; i < rc; i++ {
		r += c
	}

	return r
}
