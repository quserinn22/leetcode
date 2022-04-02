func arraySign(nums []int) int {
    ret := 1
    for _, num := range nums {
        if num == 0 {
            return 0
        }
        if num < 0 {
            ret *= -1
        }
    }
    return ret
}