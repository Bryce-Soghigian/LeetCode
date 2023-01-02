/**
 * @param {number[]} nums
 * @return {number[]}
 */
var smallerNumbersThanCurrent = function(nums) {

//FOr each number we want to go through sorted and stop when we reach the correct number
let output = []
for(let i= 0;i<nums.length;i++){
    let total = 0
    for(let j = 0;j<nums.length;j++){
        if(nums[j] < nums[i]){
        total++
        }
    }
    output.push(total)
}
   return output
};