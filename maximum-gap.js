/**
 * @param {number[]} nums
 * @return {number}
 */
var maximumGap = function(nums) {
    if(nums.length < 2) return 0
    let maxGap = -Infinity
    nums.sort((a,b) => b-a)
    
    for(let i = 0;i<nums.length-1;i++){
        const gap = nums[i] - nums[i+1]
        maxGap = Math.max(maxGap,gap)
    }
    return maxGap
    
    
};