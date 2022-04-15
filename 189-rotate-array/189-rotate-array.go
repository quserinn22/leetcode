func rotate(nums []int, k int)  {
    k %= len(nums)
    
    cache := []int{}
    for i := len(nums)-k; i < len(nums); i++ {
        cache = append(cache, nums[i])
    }
    for i := len(nums)-1; i >= k; i-- {
        nums[i] = nums[i-k]
    }
    for i := 0; i < k; i++ {
        nums[i] = cache[i]
    }
}