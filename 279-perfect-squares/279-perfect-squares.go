func dp(nums []int, target int, cache []int) int{
    if target < 0 {
        return 0
    }
    if target == 0 {
        return 0
    }
    
    if cache[target] != -1 {
        return cache[target]
    }
    
    ret := 10001
    for _, num := range nums {
        if target - num >= 0 {
            ret = min(ret, dp(nums, target - num, cache) + 1)
        } else {
            break
        }
    }
    
    cache[target] = ret
    return ret
}


func numSquares(n int) int {
    var squares []int
    for i := 1; i <= n; i++ {
        if i * i > n {
            break
        }
        squares = append(squares, i * i)
    }
    
    cache := make([]int, n+1)
    for i := range cache {
        cache[i] = -1
    }
    return dp(squares, n, cache)
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}