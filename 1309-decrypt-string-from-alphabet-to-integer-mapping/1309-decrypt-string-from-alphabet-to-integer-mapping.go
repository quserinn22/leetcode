func freqAlphabets(s string) string {
    ret := ""
    for i := 0; i < len(s); i++ {
        if s[i] == '1' || s[i] == '2' {
            if i + 2 < len(s) && s[i+2] == '#'  {
                if s[i] == '1' {
                    ret += string((s[i+1] - '1') + 'a' + 10)
                } else {
                    ret += string((s[i+1] - '1') + 'a' + 20)
                }
                i += 2
            } else {
                ret += string(s[i] - '1' + 'a')
            }
        } else {
            ret += string(s[i] - '1' + 'a')
        }
    }
    return ret
}