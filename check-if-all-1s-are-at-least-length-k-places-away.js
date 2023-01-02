/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var kLengthApart = function(nums, k) {
let count = k;
    
    for(let i = 0;i<nums.length;i++){
        let num = nums[i]
        if(num === 1){
            if(count < k){
                return false
            }
            count = 0
        }else{
          count += 1  
        }
    }
    return true
};