func reverseWords(s string) string {
    strList := []string{}
    tmp := 0
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            strList = append(strList, s[tmp:i])
            tmp = i+1
        }
    }
    strList = append(strList, s[tmp:])
    
    ret := ""
    for _, s := range strList {
        ret += " "
        ret += reverseWord(s)
    }
    
    return ret[1:]
}

func reverseWord(s string) string {
    ret := []byte{}
    
    for i := 0; i < len(s); i++ {
        ret = append(ret, s[len(s) - i - 1])
    }
    
    return string(ret)
}