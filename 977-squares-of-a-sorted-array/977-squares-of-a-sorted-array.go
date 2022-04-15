func sortedSquares(nums []int) []int {
    mid := -1
    for i := 0; i < len(nums); i++ {
        if nums[i] >= 0 {
            mid = i
            break
        }
    }
    
    if mid == -1 {
        mid = len(nums)
    }
    
    left := mid - 1
    right := mid
    
    ret := []int{}
    for left >= 0 || right < len(nums) {
        if left == -1 {
            ret = append(ret, nums[right] * nums[right])
            right += 1
            continue
        } 
        
        if right == len(nums) {
            ret = append(ret, nums[left] * nums[left])
            left -= 1
            continue
        }
        
        a := nums[right] * nums[right]
        b := nums[left] * nums[left]
        
        if a > b {
            ret = append(ret, b)
            left -= 1
        } else {
            ret = append(ret, a)
            right += 1
        }
    }
    
    return ret
}