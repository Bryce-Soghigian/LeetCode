func sum(nums []int) int{
    output := 0;
    
    for i := 0;i< len(nums); i++ {
        output += nums[i];

    };
    return output;
}

func missingNumber(nums []int) int {
    totalGaussianEstimatedValue := len(nums) * (len(nums) + 1) / 2;
    var targetNumber int = totalGaussianEstimatedValue - sum(nums);
    return targetNumber;
}