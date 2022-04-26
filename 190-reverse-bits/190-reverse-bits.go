func reverseBits(num uint32) uint32 {
    var ret uint32
    for i := 0; i < 31; i++ {
        ret += num % 2
        num /= 2
        ret *= 2
    }
    
    ret += num % 2
    
    return ret
}