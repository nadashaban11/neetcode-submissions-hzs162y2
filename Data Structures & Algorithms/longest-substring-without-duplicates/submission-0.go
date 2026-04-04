func lengthOfLongestSubstring(s string) int {
	longest := 0
	isSeen := make(map[byte]bool)
	l := 0

	for r := 0; r < len(s); r++ {

		for isSeen[s[r]] {
			isSeen[s[l]] = false
			l++
		}

		isSeen[s[r]] = true
		longest = max(longest, r - l + 1)
	}

	return longest
}