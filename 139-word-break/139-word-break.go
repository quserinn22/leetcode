
func dp(s string, cache map[string]bool, wordDict []string) bool {
    if s == "" {
        return true
    }
    if chk, ok := cache[s]; ok {
        return chk
    }
    
    for _, word := range wordDict {
        for i := 0; i <= len(s) - len(word); i++ {
            if s[i:i+len(word)] == word {
                chk := true
                if i != len(s) - len(word) {
                    chk = dp(s[i+len(word):], cache, wordDict)
                }
                if chk && dp(s[:i], cache, wordDict){
                    cache[s] = true
                    return true
                }
            }
        }
    }
    cache[s] = false
    return false
}


func wordBreak(s string, wordDict []string) bool {
    cache := map[string]bool{}
    ret := dp(s, cache, wordDict)
    return ret
}