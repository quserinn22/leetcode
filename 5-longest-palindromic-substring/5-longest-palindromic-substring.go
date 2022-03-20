func longestPalindrome(s string) string {
    n := len(s)
    cache := make([][]bool, n)
    for i := range cache {
        cache[i] = make([]bool, n)
    }
    
    ans := s[:1]
    for i := range s {
        cache[i][i] = true
        if i != n-1 {
            if s[i] == s[i+1] {
                cache[i][i+1] = true
                ans = s[i:i+2]
            }
        }
    }
    
    for i := 2; i <= n; i++ {
        for j := 0; j < n - i; j++ {
            if s[j] == s[j + i] {
                if cache[j+1][j+i-1] {
                    cache[j][j+i] = true
                    ans = s[j:j+i+1]
                }
            }
        }
    }
    
    return ans
}