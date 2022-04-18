func lengthOfLongestSubstring(s string) int {
    cache := map[rune]int{}
    
    ret := 0
    cnt := 0
    lastIdx := 0
    for i, c := range s {
        if idx, ok := cache[c]; ok {
            if idx < lastIdx {
                cache[c] = i
                cnt += 1
                continue
            }
            ret = max(ret, cnt)
            cnt = cnt - (idx - lastIdx)
            lastIdx = idx + 1
            cache[c] = i
        } else {
            cache[c] = i
            cnt += 1
        }
    }
    
    ret = max(ret, cnt)
    
    return ret
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}