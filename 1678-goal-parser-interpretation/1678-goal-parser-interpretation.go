func interpret(command string) string {
    ret := ""
    i := 0
    for i < len(command){
        if command[i] == 'G' {
            ret += "G"
            i += 1
            continue
        } else {
            if command[i+1] == ')' {
                ret += "o"
                i += 2
            } else {
                ret += "al"
                i += 4
            }
        }
    }
    return ret
}