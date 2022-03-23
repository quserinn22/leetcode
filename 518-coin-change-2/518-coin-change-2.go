func change(amount int, coins []int) int {
    if amount == 0 {
        return 1
    }
    cache := make([]int, 5001)
    for _, coin := range coins {
        cache[coin] += 1
        for i := coin; i <= amount; i+=1{
            cache[i] += cache[i-coin]
        }
    }
    
    return cache[amount]
}