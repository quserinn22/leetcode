func dp(s string, t string, cache [][]int, si int, ti int) bool {
    if si >= len(s) {
        return true
    }
    
    if ti >= len(t) {
        return false
    }
    
    if cache[si][ti] != 0 {
        if cache[si][ti] == 1 {
            return true
        } else {
            return false
        }
    }
    
    for i := ti; i < len(t); i++ {
        if t[i] == s[si] {
            if dp(s, t, cache, si+1, i+1) {
                cache[si][ti] = 1
                return true
            }
        }
    }
    cache[si][ti] = -1
    return false
}


func isSubsequence(s string, t string) bool {
    n := len(s)
    m := len(t)
    
    cache := make([][]int, n)
    for i := range cache {
        cache[i] = make([]int, m)
    }
    
    return dp(s, t, cache, 0, 0)
}