func isAlienSorted(words []string, order string) bool {
    orderMap := map[byte]int{}
    for i := 0; i < len(order); i++ {
        orderMap[order[i]] = i
    }
    
    for i := 0; i < len(words) - 1; i++ {
        if !isSmaller(words[i], words[i+1], orderMap) {
            return false
        }
    }
    return true
}

func isSmaller(a, b string, orderMap map[byte]int) bool {
    for i := 0; i < min(len(a), len(b)); i++ {
        if a[i] == b[i] {
            continue
        }
        return orderMap[a[i]] < orderMap[b[i]]
    }
    
    return len(a) <= len(b)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}