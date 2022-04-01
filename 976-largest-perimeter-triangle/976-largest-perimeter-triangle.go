func largestPerimeter(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    
    for i := len(nums) -1; i >= 2; i-- {
        for j := i-1; j >= 1; j-- {
            for k := j-1; k >= 0; k-- {
                if nums[i] < nums[k] + nums[j] {
                    return nums[k] + nums[i] + nums[j]
                }
            }
        }
    }
    return 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}