func hammingWeight(num uint32) int {
    
    cnt := 0
    for num > 0 {
        cnt += int(num % 2)
        num /= 2
    }
    return cnt
}