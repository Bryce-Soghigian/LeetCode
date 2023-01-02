/**
 * @param {number[]} nums
 * @return {number}
 */
var findLengthOfLCIS = function(nums) {
    if(nums.length === 0) return 0
    let count = 0;
    let curr = 0;
    for(let i = 0;i<nums.length;i++){
        if(nums[i] < nums[i+1]){
            curr++
            if(curr > count){
                count = curr
            }
        }else{
            curr = 0
        }
    }
    return count +1
};