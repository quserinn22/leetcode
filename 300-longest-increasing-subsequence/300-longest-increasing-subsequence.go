func dp(nums []int, cache []int, idx int) int {
    if cache[idx] != 0 {
        return cache[idx]
    }
    
    cache[idx] = 1
    for i := idx-1; i >= 0; i--{
        if nums[idx] > nums[i] {
            cache[idx] = max(cache[idx], dp(nums, cache, i) + 1)
        }
    }   

    return cache[idx]
}



func lengthOfLIS(nums []int) int {
    cache := make([]int, len(nums))
    
    ans := 1
    for i := range nums {
        ans = max(ans, dp(nums, cache, i))
    }
     
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
