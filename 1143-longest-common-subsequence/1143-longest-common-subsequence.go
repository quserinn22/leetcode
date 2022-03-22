func dp(s string, t string, cache [][]int, si int, ti int) int {
    if si >= len(s) {
        return 0
    }
    
    if ti >= len(t) {
        return 0
    }
    
    if cache[si][ti] != -1 {
        if cache[si][ti] == -2 {
            return -1
        }
        return cache[si][ti]
    }
    
    ret := -2
    for i := ti; i < len(t); i++ {
        if t[i] == s[si] {
            if dp(s, t, cache, si+1, i+1) != -1 {
                ret = max(ret, dp(s, t, cache, si+1, i+1) + 1) 
            }
        }
    }
        
    if si < len(s) {
        if dp(s, t, cache, si+1, ti) != -1 {
            ret = max(ret, dp(s, t, cache, si+1, ti))
        }
    }
        
    cache[si][ti] = ret
    
    return cache[si][ti]
}


func longestCommonSubsequence(text1 string, text2 string) int {
    n := len(text1) 
    m := len(text2)
    t1 := text1
    t2 := text2
    
    if n < m {
        t1 = text2
        t2 = text1
        n = m
        m = len(t2)
    }
    
    cache := make([][]int, m)
    for i := range cache {
        cache[i] = make([]int, n)
        for j := range cache[i] {
            cache[i][j] = -1
        }
    }
    
    return dp(t2, t1, cache, 0, 0)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}