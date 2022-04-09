func sortByBits(arr []int) []int {
    sort.Slice(arr, func(i, j int) bool {
        a := arr[i]
        b := arr[j]
        a1 := 0
        b1 := 0
        for a > 0 {
            a1 += a %2
            a /= 2
        }
        for b > 0 {
            b1 += b % 2
            b /= 2
        }
        
        if a1 == b1 {
            return arr[i] < arr[j]
        }
        
        return a1 < b1
    })
    
    return arr
}