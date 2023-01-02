func halvesAreAlike(s string) bool {
    var vowelsCount int
    var n int = len(s)
	vowels := map[string]bool{"a":true, "i":true, "u":true, "e":true, "o":true, "A":true, "I":true, "U":true, "E":true, "O":true}
    for i := 0; i < n / 2; i++ {
        if ok := vowels[string(s[i])]; ok {
            vowelsCount++
        }
    }
    for i := n / 2; i < n; i++ {
        if ok := vowels[string(s[i])]; ok {
            vowelsCount--
        }
    }

	return vowelsCount == 0
}