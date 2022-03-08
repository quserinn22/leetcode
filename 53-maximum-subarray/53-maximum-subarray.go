var cache []int

func dp(nums []int, idx int) int {
    if cache[idx] != -1 {
        return cache[idx]
    }
    
    if idx == 0 {
        return max(0, nums[0])
    }
    
    cache[idx] = max(dp(nums, idx - 1) + nums[idx], 0)
    return cache[idx]
}

func maxSubArray(nums []int) int {
    cache = make([]int, 100000)
    for i, _ := range nums {
        cache[i] = -1
    }
    
    cache[0] = max(0, nums[0])
    
    dp(nums, len(nums) - 1)
    ret := 0
    for i, c := range cache {
        if i == len(nums) {
            break
        }
        if c > ret {
            ret = c
        }
    }
    
    if ret == 0 {
        ret = -10000
        for _, num := range nums {
            if num > ret {
                ret = num
            }
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