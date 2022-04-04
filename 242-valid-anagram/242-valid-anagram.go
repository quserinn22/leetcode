func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    cache := make(map[rune]int)
    for _, r := range s {
        if _, ok := cache[r]; ok {
            cache[r] += 1
        } else {
            cache[r] = 1
        }
    }
    
    for _, r := range t {
        if _, ok := cache[r]; ok {
            cache[r] -= 1
            if cache[r] < 0 {
                return false
            }
        } else {
            return false
        }
    }
    
    return true
}