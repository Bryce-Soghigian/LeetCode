/**
 * @param {number[]} heights
 * @return {number}
 */
var heightChecker = function(heights) {
    sorted = heights.slice().sort((a,b)=>a-b);
    let counter = 0;
    for(let i in heights){
        if(heights[i] - sorted[i]) counter++;
    }
    
    return counter;
};