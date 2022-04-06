func findTheDifference(s string, t string) byte {
    cache := map[rune]int{}
    
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
                return byte(r)
            }
        } else {
            return byte(r)
        }
    }
    
    return byte('a')
}