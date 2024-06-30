package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string
	for _, r := range s {
		if r == ' ' || r == '\n' || r == '\t' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word = word + string(r)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}

func PrintWordsTables(a []string) {
	for _, word := range a {
		println(word)
	}
}
