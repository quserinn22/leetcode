
func maxScoreSightseeingPair(values []int) int {
    rightMax := values[len(values) - 1] - (len(values) - 1)
    ret := 0
    for i := len(values)-2; i >=0; i-- {
        rightMax = max(rightMax, values[i+1]-(i+1))
        ret = max(values[i] + i + rightMax, ret)
    }
    return ret
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}