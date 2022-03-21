func dp(nums []int, cache [][]int, idx int, chk int) int {
    if cache[idx][chk] != 0 {
        return cache[idx][chk]
    }
    
    cache[idx][chk] = 1
    for i := idx -1; i >= 0; i-- {
        if chk == 0 && nums[i] > nums[idx] {
            cache[idx][chk] = max(cache[idx][chk], dp(nums, cache, i, 1) + 1)
        }
        if chk == 1 && nums[i] < nums[idx] {
            cache[idx][chk] = max(cache[idx][chk], dp(nums, cache, i, 0) + 1)
        }
    }
    
    return cache[idx][chk]
}


func wiggleMaxLength(nums []int) int {
    n := len(nums)
    cache := make([][]int, n)
    for i := range cache {
        cache[i] = make([]int, 2)
    }
    
    ans := 1
    for i := range nums {
        ans = max(ans, dp(nums, cache, i, 0))
        ans = max(ans, dp(nums, cache, i, 1))
    }
    
    return ans
}

func max(a, b int) int {
    if a > b  {
        return a
    }
    return b
}