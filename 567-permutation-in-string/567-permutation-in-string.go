func checkInclusion(s1 string, s2 string) bool {
    cache := map[byte]int{}
    for i := range s1 {
        cache[s1[i]] += 1
    }
    
    for i := 0; i < len(s2); i++ {
        c := s2[i]
        if _, ok := cache[c]; ok {
            cache[c] -= 1
        } else {
            cache[c] = -1
        }
        
        if i - len(s1) >= 0 {
            cache[s2[i-len(s1)]] += 1
        }
        
        chk := true
        for _, v := range cache {
            if v != 0 {
                chk = false
                break
            }
        }
        if chk {
            return true
        }
    }
    return false
}