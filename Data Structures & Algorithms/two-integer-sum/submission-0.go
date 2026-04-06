func twoSum(nums []int, target int) []int {
    exist := make(map[int]int)

	for i:= 0; i < len(nums); i++ {
		ans := target - nums[i]
		if idx, ok := exist[ans]; ok {
			if i < idx {
				return []int{i, idx}
			}
			return []int{idx, i}
		}
		
		exist[nums[i]] = i 
	}
	return nil
}

// Time Complexity = O(n)
// Space Complexity = O(n)
