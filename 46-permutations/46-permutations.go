func permute(nums []int) [][]int {
    ret := [][]int{}
    
    var backtrace func(int)
    backtrace = func(start int) {
        if start == len(nums) - 1 {
            ans := make([]int, len(nums))
            copy(ans, nums)
            ret = append(ret, ans)
            return 
        }
        for i := start; i < len(nums); i++ {
            nums[i], nums[start] = nums[start], nums[i]
            backtrace(start + 1)
            nums[i], nums[start] = nums[start], nums[i]
        }
    }
    
    backtrace(0)
    
    return ret
}

