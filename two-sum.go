func twoSum(nums []int, target int) []int {
    output := []int{-1, -1}
    complements := map[int]int{}
    
    for i := 0; i < len(nums); i++ {
        targetPair := target - nums[i]
        if complementIndex, ok := complements[targetPair]; ok {
            output[0] = complementIndex
            output[1] = i
            return output
        } 
        complements[nums[i]] = i
    }
    return output
    
}