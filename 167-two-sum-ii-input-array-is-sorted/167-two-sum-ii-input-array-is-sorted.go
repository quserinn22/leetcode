func twoSum(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++ {
        tmp := target - numbers[i]
        
        left := i+1
        right := len(numbers) - 1
        
        for left <= right {
            mid := (left + right) / 2
            
            if numbers[mid] > tmp {
                right = mid - 1
            } else if numbers[mid] < tmp {
                left = mid + 1
            } else {
                return []int{i+1, mid+1}
            }
        }
    }
    
    return []int{0, 0}
}