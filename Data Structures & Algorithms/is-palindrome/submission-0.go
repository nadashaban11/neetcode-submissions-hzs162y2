func isPalindrome(s string) bool {
	j := len(s) - 1;
	for i:= 0; i <= j; {
		if i >= j {
			return true
		}
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
			continue
		}
		if unicode.ToUpper(rune(s[i])) != unicode.ToUpper(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
