func nthUglyNumber(n int) int {
    i2 := 0
    i3 := 0
    i5 := 0
    
    uglyNums := []int{1}
    cnt := 0
    for {
        cnt += 1
        if cnt == n {
            return uglyNums[len(uglyNums) - 1]
        }
        
        n1 := uglyNums[i2] * 2
        n2 := uglyNums[i3] * 3
        n3 := uglyNums[i5] * 5
     
        n := min(min(n1, n2), n3)
        if n == n1 {
            i2 += 1
        }
        if n == n2 {
            i3 += 1
        }
        if n == n3 {
            i5 += 1
        }
        uglyNums = append(uglyNums, n)
    }
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}