package piscine

func ConcatParams(args []string) string {
	var result string
	for i, c := range args {
		if i == 0 {
			result = c
		} else {
			result = result + "\n" + c
		}
	}
	return result
}
