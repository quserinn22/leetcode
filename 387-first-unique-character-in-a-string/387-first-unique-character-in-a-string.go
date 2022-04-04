func firstUniqChar(s string) int {
    cache := make(map[rune]int)
    for _, c := range s {
        if _, ok := cache[c]; ok {
            cache[c] += 1
        } else {
            cache[c] = 1
        }
    }
    
    for i, c := range s {
        if cache[c] == 1 {
            return i
        }
    }
    return -1
}