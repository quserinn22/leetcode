func average(salary []int) float64 {
    sum := 0
    max := salary[0]
    min := salary[0]
    for _, s := range salary {
        sum += s
        if s > max {
            max = s
        }
        if s < min {
            min = s
        }
    }
    return float64(sum - max - min) / float64(len(salary) - 2)
}