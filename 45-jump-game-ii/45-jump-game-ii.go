var cache []int // -1: false, 1: true


func dp(nums []int, idx int) int {
    if cache[idx] != 10000 {
        return cache[idx]
    } 
    
    for i := idx - 1; i >= 0; i-- {
        if idx - nums[i] <= i {
            cnt := dp(nums, i)
            if cnt < cache[idx] {
                cache[idx] = cnt + 1
            }
        }
    }
    
    return cache[idx]
}

func jump(nums []int) int {
    cache = make([]int, 10000)
    cache[0] = 0
    for i := 1; i < 10000; i++ {
        cache[i] = 10000
    }
    
    return dp(nums, len(nums) - 1)
}