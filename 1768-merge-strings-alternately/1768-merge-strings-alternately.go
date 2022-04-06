func mergeAlternately(word1 string, word2 string) string {
    idx1 := 0
    idx2 := 0
    ret := ""
    for idx1 < len(word1) || idx2 < len(word2) {
        if idx1 == len(word1) {
            ret += word2[idx2:]
            return ret
        } else if idx2 == len(word2) {
            ret += word1[idx1:]
            return ret
        } else {
            if idx1 <= idx2 {
                ret += string(word1[idx1])
                idx1 += 1
            } else {
                ret += string(word2[idx2])
                idx2 += 1
            }
        }
    }
    
    return ret
}