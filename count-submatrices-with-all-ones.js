/**
 * @param {number[][]} mat
 * @return {number}
 */
var numSubmat = function(matrix) {
  if (matrix.length <= 0) 
      return 0;
  const m = matrix.length;
  const n = matrix[0].length;
  const dp = new Array(m).fill(null).map(() => new Array(n).fill(0));
  // 1. build preSum on each row: dp[][]
   for(var i=0; i<m; i++)
      for(var j=0; j<n; j++)
      {
        dp[i][j] =  matrix[i][j]===0 ? 0 : (j===0? matrix[i][0]: dp[i][j-1] + 1);   
      }
    
  // 2. count all from reversed-column order
  var result = 0;
  for (var i = 0; i<m; i++) {
    for (var j = n-1; j >=0 ; j--) {
      if (matrix[i][j] === 0)
          continue;
      var min = dp[i][j];
      for (var k = i; k <m; k++) {
        if (dp[k][j] === 0) 
            break;
        min = Math.min(dp[k][j], min);
        result += min;
      }
    }
  }
  return result;
};