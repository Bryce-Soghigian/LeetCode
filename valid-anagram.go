func frequencyMap(s string)map[rune]int{
    output := make(map[rune]int, 26)
    for _, key := range s {
        if _, ok := output[key]; ok {
            output[key] += 1
        } else { 
            output[key] = 1
        }
    }
    return output
}

func isSameMap(m1 map[rune]int, m2 map[rune]int) bool {
    for key, value := range m1 {
        other, ok := m2[key]

        if !ok || other != value {
            return false
        }
    }
    return true
}
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m1 := frequencyMap(s)
    m2 := frequencyMap(t)
    return isSameMap(m1, m2)
}