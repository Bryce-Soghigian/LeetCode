/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxSubArrayLen = function(nums, k) {
    var id = {};
    id[0] = -1;
    var sum = 0, max = 0;
    for(var i = 0; i < nums.length; i++) {
        sum += nums[i];
        if (!(sum in id)) {
            id[sum] = i;
        }
        if(sum - k in id) {
            max = Math.max(max, i - id[sum - k]);
        }    
    }
    return max;
};