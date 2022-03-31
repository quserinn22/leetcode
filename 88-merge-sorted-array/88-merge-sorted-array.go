func merge(nums1 []int, m int, nums2 []int, n int) {
    i := 0
    j := 0
    
    ret := []int{}
    for i < m || j < n {
        if i == m {
            ret = append(ret, nums2[j])
            j +=1
        } else if j == n {
            ret = append(ret, nums1[i])
            i += 1
        } else {
            if nums1[i] > nums2[j] {
                ret = append(ret, nums2[j])
                j += 1
            } else {
                ret = append(ret, nums1[i])
                i += 1
            }
        }
    }
    
    for i := range ret {
        nums1[i] = ret[i]
    }
}