var cache []int // -1: false, 1: true


func dp(nums []int, idx int) bool{
    if cache[idx] == 1 {
        return true
    } else if cache[idx] == -1 {
        return false
    }
    
    for i := idx - 1; i >= 0; i-- {
        if idx - nums[i] <= i && dp(nums, i) {
            cache[idx] = 1
            return true
        }
    }
    
    cache[idx] = -1
    return false
}

func canJump(nums []int) bool {
    cache = make([]int, 10000)
    cache[0] = 1
    return dp(nums, len(nums) - 1)
}