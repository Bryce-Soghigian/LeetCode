/**
 * @param {number} n - a positive integer
 * @return {number}
 */
var hammingWeight = function(n) {
    let count;
 
    for (count=0; n; count++)
        n &= n - 1;
    return count;
};