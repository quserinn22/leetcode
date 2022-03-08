func dp(nums []int, idx int, cache []int) int {
    if cache[idx] != -1 {
        return cache[idx]
    }
    
    if idx == 0 {
        return max(0, nums[0])
    }
    
    cache[idx] = max(dp(nums, idx - 1, cache) + nums[idx], 0)
    return cache[idx]
}

func dp2(nums []int, idx int, cache []int) int {
    if cache[idx] != 1 {
        return cache[idx]
    }
    
    if idx == 0 {
        return min(0, nums[0])
    }
    
    cache[idx] = min(dp2(nums, idx - 1, cache) + nums[idx], 0)
    return cache[idx]
}

func find(cache []int, nums []int) int {
    ret := 0
    for _, c := range cache {
        ret = max(c, ret)
    }

    if ret == 0 {
        ret = -30000
        for _, num := range nums {
            ret = max(num, ret)
        }
    }
    return ret
}

func find2(cache []int, nums []int) int {
    ret := 0
    for _, c := range cache {
        ret = min(c, ret)
    }
    
    if ret == 0 {
        ret = 30000
        for _, num := range nums {
            ret = min(num, ret)
        }
    }
    return ret
}

func maxSubarraySumCircular(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    cache := make([]int, len(nums))
    for i := range nums {
        cache[i] = -1
    }
    dp(nums, len(nums) - 1, cache)
    maxSub := find(cache, nums)

    for i := range nums {
        cache[i] = 1
    }
    dp2(nums, len(nums) - 1, cache)
    minSub := find2(cache, nums)
    
    sum := 0
    for _, num := range nums {
        sum += num
    }
        
    
    ret := max(maxSub, sum - minSub)
    if ret == 0 {
        ret = -30000
        for _, num := range nums {
            ret = max(ret, num)
        }
    }
    
    return ret
}

func max(a, b int) int{
    if a > b {
        return a
    }
    
    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }
    
    return a
}