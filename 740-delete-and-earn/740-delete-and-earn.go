var cache []int
var cached []bool

func max(a, b int) int{
    if a > b {
        return a
    }
    
    return b
}

func dp(nums []int, idx int) int{
    if idx < 2 {
        return nums[idx]
    }
    
    if idx == len(nums) {
        return max(dp(nums, idx - 1), dp(nums, idx - 2))
    }
    
    if cached[idx] {
        return cache[idx]
    }
    
    if idx == 2 {
        cache[2] = nums[0] + nums[2]
        cached[2] = true
        return cache[2]
    }
    
    cache[idx] = max(dp(nums, idx-2) + nums[idx], dp(nums, idx-3) + nums[idx])
    cached[idx] = true
    return cache[idx]
}


func deleteAndEarn(nums []int) int {
    earns := make([]int, 10001)
    for _, num := range nums {
        earns[num] += num
    }
    
    cache = make([]int, 10001)
    cached = make([]bool, 10001)
    
    return dp(earns, 10001)
}