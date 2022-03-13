func numberOfArithmeticSlices(nums []int) int {
    if len(nums) < 3 {
        return 0
    }
    
    cache := make([]map[int]bool, len(nums))
    for i := range cache {
        cache[i] = map[int]bool{}
    }
    
    cnt := 0
    var diff1, diff2 int 
    for i := range nums[:len(nums)-2] {
        j := i+1
        k := j+1
        diff1 = nums[j] - nums[i]
        diff2 = nums[k] - nums[j]
        if diff1 != diff2 {
            continue
        }

        if cache[i][diff1] {
            continue
        }
        
        cache[i][diff1] = true
        cache[j][diff1] = true

        tempCnt := 1 
        cnt += tempCnt
        l := k + 1
        for {
            cache[l-1][diff1] = true
            if l >= len(nums) || nums[l] - nums[l-1] != diff2 {
                break
            }
            tempCnt +=1 
            cnt += tempCnt
            l += 1
        }

    }
    return cnt
}