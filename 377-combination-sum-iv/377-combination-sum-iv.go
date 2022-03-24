func dp(nums []int, target int, cache []int) int{
    if target < 0 {
        return 0
    }
    if target == 0 {
        return 1
    }
    
    if cache[target] != -1 {
        return cache[target]
    }
    
    ret := 0
    for _, num := range nums {
        ret += dp(nums, target - num, cache)
    }
    cache[target] = ret
    return ret
}

func combinationSum4(nums []int, target int) int {
    cache := make([]int, target + 1)
    for i := range cache {
        cache[i] = -1
    }
    return dp(nums, target, cache)
}