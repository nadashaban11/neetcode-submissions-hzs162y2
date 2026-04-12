func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return 0
	}
	cnt := 1
	ans := 1
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		if nums[i+1] - nums[i] == 1 {
			cnt++
		} else{
			ans = max(ans, cnt)
			cnt = 1
		}
	}
	return max(ans, cnt)
}

// Time Complexity = O(nlogn)
