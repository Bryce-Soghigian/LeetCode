/**
 * @param {number[]} nums
 * @return {number}
 */
var findLHS = function(nums) {
    let dict = {}
    let res = 0
    for(let i of nums){
    dict[i] = dict[i] + 1 || 1
    }
    for(let key in dict){
        if(dict[parseInt(key) + 1]){
            res = Math.max(res, dict[key] + dict[parseInt(key) + 1])
        }
    }
    
    
    return res
};