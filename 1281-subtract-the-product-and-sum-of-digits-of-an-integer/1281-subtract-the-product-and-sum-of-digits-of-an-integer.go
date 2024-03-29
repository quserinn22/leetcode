func subtractProductAndSum(n int) int {
    prod := 1
    sum := 0
    
    for n > 0 {
        prod *= n % 10 
        sum += n % 10
        n /= 10
    }
    
    return prod - sum
}