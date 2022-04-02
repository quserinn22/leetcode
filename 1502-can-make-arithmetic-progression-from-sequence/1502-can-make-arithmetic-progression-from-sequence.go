func canMakeArithmeticProgression(arr []int) bool {
    sort.Slice(arr, func(a, b int) bool {
        return arr[a] < arr[b]
    })
    
    diff := arr[1] - arr[0]
    for i := 1; i < len(arr) - 1; i++ {
        if arr[i+1] - arr[i] != diff {
            return false
        }
    }
    return true
}