/**
 * @param {number[]} candyType
 * @return {number}
 */
var distributeCandies = function(candyType) {
    //We want to count all of the candies nd then we wnt to return 
    
    const max = candyType.length / 2;
    //Max num of candy
    
    
    const set = new Set();
    for(let i = 0;i<candyType.length;i++){
        set.add(candyType[i])
    }
    
    
    if(set.size < max) return set.size
    else return max
};