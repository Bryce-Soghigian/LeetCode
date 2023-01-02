/**
 * @param {number[]} nums
 * @return {number}
 */
var thirdMax = function(nums) {    
    let max = new Set()
    
    for (let i in nums) {
        max.add(nums[i])
		// maintain atmost 3 elements by removing min
        if (max.size > 3) {
            let m = Array.from(max)
            max.delete(Math.min(...m))
        }
    }
    
    let m = Array.from(max)
    return m.length === 3 ? Math.min(...m) : Math.max(...m)
    //  nums = Array.from(new Set(nums.sort((a,b) => a-b)))
    // if(nums.length < 3){
    //     return nums[nums.length-1]
    // }
    // return nums[nums.length -3]
};