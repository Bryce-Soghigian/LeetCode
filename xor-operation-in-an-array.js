/**
 * @param {number} n
 * @param {number} start
 * @return {number}
 */
var xorOperation = function(n, start) {
  let result = start;
  
  for (let i = 1; i < n; i ++) {
    result ^= start + i * 2;
  }
  
  return result;
};