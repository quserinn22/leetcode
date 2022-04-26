func isPowerOfTwo(n int) bool {
    p := 1
    
    if n <= 0 {
        return false
    }
    
    for p < n {
        if n % p != 0 {
            return false
        }
        
        p *= 2
    }
    
    return true
}