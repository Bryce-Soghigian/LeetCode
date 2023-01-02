func uniqueOccurrences(arr []int) bool {
    seen := map[int]int{}
    for _, val := range arr {
        seen[val] += 1
    }

    unique := map[int]bool{}
    for _, freq := range seen {
        if _, ok := unique[freq]; ok {
            return false
        }
        unique[freq] = true
    }
    return true
}