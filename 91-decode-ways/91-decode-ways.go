func dp(s string, idx int, cache []int) int {
    if idx > len(s) - 1 {
        return 0
    }
    
    if s[idx] == '0' {
        return 0
    }
    
    if idx == len(s) - 1 {
        return 1
    }
    
    if cache[idx] != -1 {
        return cache[idx]
    }
    
    cache[idx] = dp(s, idx+1, cache)
    if s[idx] == '1' || (s[idx] == '2' && s[idx+1] < '7'){
        if idx + 2 > len(s) - 1 {
            cache[idx] += 1
        } else {
            cache[idx] += dp(s, idx+2, cache)
        }
    }
    return cache[idx]
}

func numDecodings(s string) int {
    if s == "0" {
        return 0
    }
    
    cache := make([]int, len(s))
    for i := range cache {
        cache[i] = -1
    }
    return dp(s, 0, cache)
}