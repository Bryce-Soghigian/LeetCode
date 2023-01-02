/**
 * @param {number} N
 * @return {number}
 */
var maxA = function(N) {
   /**

[1, ... n]
for(let )


let dp = new Array(N +1).fill(0)
for(let i =1;i<= N;i++){

}

   **/
    
    let dp = new Array(N +1).fill(0)
    dp[0] = 0;
    for(let i = 1;i<= N;i++){
        dp[i] = i
        for(let j = 3;j<i;j++){
             dp[i]  = Math.max(dp[i],dp[i-j] * (j-1))
        }
    }
   return dp[N]
}