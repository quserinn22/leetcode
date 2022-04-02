func areAlmostEqual(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    
    n := len(s1)
    
    cnt := 0
    var a, b, c, d string
    for i := 0; i < n; i++ {
        if s1[i] != s2[i] {
            cnt += 1
            if a == "" {
                a = string(s1[i])
                b = string(s2[i])
            } else {
                c = string(s1[i])
                d = string(s2[i])
            }
            
        }
        if cnt > 2 {
            return false
        }
    }
    
    if a == d && b == c {
        return true
    }
    return false
}