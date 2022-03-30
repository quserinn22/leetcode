func countOdds(low int, high int) int {
    cnt := (high - low) / 2
    if high % 2 == 0 && low %2 == 0 {
        return cnt
    } else {
        return cnt + 1
    }
}