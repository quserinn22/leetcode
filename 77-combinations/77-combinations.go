func combine(n int, k int) [][]int {
    cache := make([][][][]int, 21)
    
    for i := range cache {
        cache[i] = make([][][]int, 21)
    }
    
    return solve(n, k, cache) 
}

func solve(n int, k int, cache [][][][]int) [][]int {
    ret := [][]int{}
    
    if cache[n][k] != nil {
        return cache[n][k]
    }
    
    if k == 1 {
        for i := 1; i <= n; i++ {
            ret = append(ret, []int{i})
        }
        return ret
    }
    
    if k == 2 {
        for i := 1; i < n; i++ {
            for j := i+1; j <= n; j++ {
                ret = append(ret, []int{i, j})
            }
        }
        return ret
    }
    
    for i := 1; i <= n - k + 1; i++ {
        tmp := combine(n-i, k-1)
        for j := range tmp {
            tmp[j] = append(tmp[j], n-i+1)
        }
        ret = append(ret, tmp...)
    }
    
    cache[n][k] = ret
    
    return ret
}

