func nextGreaterElement(nums1 []int, nums2 []int) []int {
    
    cache := make(map[int]int)
    for i, num := range nums2 {
        cache[num] = i
    }
    
    ret := []int{}
    for _, num := range nums1 {
        chk := false
        for i := cache[num]; i < len(nums2); i++ {
            if nums2[i] > num {
                ret = append(ret, nums2[i])
                chk = true
                break
            }
        }
        if !chk {
            ret = append(ret, -1)
        }
        
    }
    
    return ret
}