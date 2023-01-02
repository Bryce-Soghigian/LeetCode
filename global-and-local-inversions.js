/**
 * @param {number[]} nums
 * @return {boolean}
 */
var isIdealPermutation = function(nums) {
    const n = nums.length;
    let localInversions = 0;
    let globalInversions = 0;
    for(let i = 0;i<n;i++){
        if(0 <= i && i< n-1 && nums[i] > nums[i+1]){
            localInversions++
        }
        for(let j = i;j<n;j++){
            /**
            
            0 <= i < j < n
nums[i] > nums[j]
*/
            if(0 <= i && i < j && j < n && nums[i] > nums[j]){
                globalInversions++
            }
        }
    }
    
    
    
    return localInversions === globalInversions
};