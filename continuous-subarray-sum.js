/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var checkSubarraySum = function(nums, k) {
    var i=0
    var sum = 0
    var sumMap = {}
    const map = new Map([[0, -1]]);
    if(nums.length <2) {
        return false
    }
    sumMap[0] = -1
    while(i<nums.length) {
        sum += nums[i]
        if(k !== 0) {
            sum = sum%k
        }
        if(sumMap[sum] != undefined) {
            if(i-sumMap[sum] >=2) {
                return true
            }
        }
        else {
            sumMap[sum] = i 
        }
        i++
    }
    return false
};