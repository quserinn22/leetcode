func letterCasePermutation(s string) []string {
    ret := []string{}
    
    s = strings.ToLower(s)
    
    char := []byte(s)
    var backtrace func([]byte, int)
    backtrace = func(ss []byte, start int) {
        if start == len(ss){
            ret = append(ret, string(ss))
            return
        }
        
        backtrace(ss, start + 1)
        
        if unicode.IsLetter(rune(ss[start])) {
            ss[start] = byte(unicode.ToUpper(rune(ss[start])))
            backtrace(ss, start + 1)
            ss[start] = byte(unicode.ToLower(rune(ss[start])))
        } 
    }
    
    backtrace(char, 0)
    
    return ret
}

