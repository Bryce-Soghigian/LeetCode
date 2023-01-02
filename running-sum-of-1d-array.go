func runningSum(nums []int) []int {
    output := make([]int, len(nums), len(nums))
    runningSum := 0
    for i := 0; i < len(nums); i++ {
        runningSum += nums[i]
        output[i] = runningSum
    }
    return output
}