type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for i := 0; i < len(strs); i++ {
		encoded += strconv.Itoa(len(strs[i])) + "#" + strs[i]
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	strs := []string {}
	i := 0
	j := i
	for i < len(encoded) {
		for encoded[j] != '#' {
			j++
		}
		size, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		strs = append(strs, encoded[i: i + size])
		i += size
		j = i
	}
	return strs
}
