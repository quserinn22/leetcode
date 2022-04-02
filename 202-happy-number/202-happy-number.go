func isHappy(n int) bool {
    cache := map[int]bool{}
    for {
        if _, ok := cache[n]; ok {
            return false
        }
        cache[n] = true
        
        n = next(n)
        if n == 1 {
            return true
        }
    }
}

func next(n int) int {
    ret := 0
    for n > 0 {
        ret += (n % 10) * (n % 10)
        n /= 10
    }
    return ret
}