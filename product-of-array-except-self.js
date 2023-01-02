/**
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function(nums) {
    let res = new Array(nums.length)
    
    for(let i = 0;i<nums.length;i++){
        let temp_value = 1
        for(let j = 0;j<nums.length;j++){
            if(j !== i){
                temp_value *= nums[j]
            }

        }
    res[i] = temp_value
    }
    return res

    
    
};