func canConstruct(ransomNote string, magazine string) bool {
    cache := make(map[rune]int)
    
    for _, s := range magazine {
        if _, ok := cache[s]; ok {
            cache[s] += 1 
        } else {
            cache[s] = 1
        }
    }
    
    for _, s := range ransomNote {
        if _, ok := cache[s]; !ok {
            return false
        } else {
            cache[s] -= 1
            if cache[s] < 0 {
                return false
            }
        }
    }
    
    return true
}