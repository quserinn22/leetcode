func maximumWealth(accounts [][]int) int {
    ret := 0
    for _, account := range accounts {
        sum := 0
        for _, money := range account {
            sum += money
        }
        if sum > ret {
            ret = sum
        }
    }
    return ret
}