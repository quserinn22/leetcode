func intersect(nums1 []int, nums2 []int) []int {
    inter := make([]int, 1001)
    
    for _, num := range nums1 {
        inter[num] += 1
    }
    
    ret := []int{}
    for _, num := range nums2 {
        if inter[num] > 0 {
            ret = append(ret, num)
            inter[num] -= 1
        }
    }
    
    return ret 
}