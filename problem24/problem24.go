package problem24

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SplitHeadAndTail(str string) (string, string) {
	currentIndex := len(str) - 1
	currentChar := ""
	for currentIndex >= 0 {
		if str[currentIndex:currentIndex+1] < currentChar {
			currentIndex++
			break
		}
		currentChar = str[currentIndex : currentIndex+1]
		currentIndex--
	}

	if currentIndex < 0 {
		return "", str
	}
	return str[0:currentIndex], str[currentIndex:len(str)]
}

// http://wordaligned.org/articles/next-permutation
func NextPermutation(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}
	head, tail := SplitHeadAndTail(str)
	if head == "" {
		return str
	}
	headChars := []rune(head)
	tailChars := []rune(tail)
	currentIndex := len(tail) - 1
	for currentIndex >= 0 {
		if tail[currentIndex:currentIndex+1] > head[len(head)-1:len(head)] {
			temp := tailChars[currentIndex]
			tailChars[currentIndex] = headChars[len(head)-1]
			headChars[len(head)-1] = temp
			break
		}
		currentIndex--
	}
	return string(headChars) + Reverse(string(tailChars))
}

func LexicographicPermutations(str string) []string {
	result := []string{str}
	last := ""
	current := NextPermutation(str)
	for last != current {
		result = append(result, current)
		last = current
		current = NextPermutation(current)
	}
	return result
}
