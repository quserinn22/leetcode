func isValid(s string) bool {
    cache := []rune{}
    for _, c := range s {
        if c == '(' || c == '{' || c == '[' {
            cache = append(cache, c)
        } else {
            if c == ')' {
                if len(cache) != 0 && cache[len(cache) - 1] == '(' {
                    cache = cache[:len(cache)-1]
                } else {
                    return false
                }
            } else if c == ']' {
                if len(cache) != 0 && cache[len(cache) - 1] == '[' {
                    cache = cache[:len(cache)-1]
                } else {
                    return false
                }
            } else {
                if len(cache) != 0 && cache[len(cache) - 1] == '{' {
                    cache = cache[:len(cache)-1]
                } else {
                    return false
                }
            }
        }
    }
    return len(cache) == 0
}