var checkPossibility = function(nums) {
    let error = 0
    let len = nums.length
    for (let i = 1; i < len; i++)
        if (nums[i] < nums[i-1]){
            if((i > 1 && i < len - 1 && nums[i-2] > nums[i] && nums[i+1] < nums[i-1]) || error++){
                return false
            }
        }
    return true
};