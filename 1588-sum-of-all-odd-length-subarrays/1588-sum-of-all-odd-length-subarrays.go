func sumOddLengthSubarrays(arr []int) int {    
    sum := 0
    
    for i := 1; i <= len(arr); i+=2 {
        for j := 0; j <= len(arr) - i; j++ {
            for k := j; k < j+i; k++ {
                sum += arr[k]
            }
        }
    }
    
    return sum
}