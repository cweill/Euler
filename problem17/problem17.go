package problem17

func NumberToWords(n int) string {
	subtwenty := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}
	tens := []string{
		"zero",
		"ten",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}
	written := ""
	n = n % 10000
	if n >= 1000 {
		written += subtwenty[n/1000] + " thousand "
	}
	n = n % 1000
	if n >= 100 {
		written += subtwenty[n/100] + " hundred "
	}
	n = n % 100

	if len(written) > 0 && n > 0 {
		written += "and "
	}

	if n >= 20 {
		written += tens[n/10] + " "
		n = n % 10
	}
	if n < 20 && n > 0 {
		written += subtwenty[n]
	}
	return written
}

func NumberLetterCount(n int) int {
	word := NumberToWords(n)
	count := 0
	for i := 0; i < len(word); i++ {
		if word[i:i+1] != " " {
			count += 1
		}
	}
	return count
}
