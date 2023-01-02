/**
 * @param {number[]} nums
 * @return {number}
 */
var maximumUniqueSubarray = function(nums) {
    let max = 0;
    let curr = 0
    let start = 0
    let seen = new Set();
    for(let end = 0; end<nums.length;end++){
        
        while(seen.has(nums[end])){
            seen.delete(nums[start])
            curr -= nums[start]
            start++
        }
        curr += nums[end]
        seen.add(nums[end])
        if(curr > max){
            max = curr
        }
    }
    return max
};