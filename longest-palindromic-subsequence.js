/**
 * @param {string} s
 * @return {number}
 */
var longestPalindromeSubseq = function(s) {
    const dp = [];
    
    for(let i = 0; i <= s.length; i++) {
        let rowElm = []
        for(let j = 0; j <= s.length; j++) {
            rowElm[j] = 0;
        }
        dp.push(rowElm);
    }
    for(let i = 1; i <= s.length; i++) {
        for(let j = 1; j <= s.length; j++) {
            let rowPos = i - 1;
            let colPos = s.length-j; //checking in reversed order
            if(s[rowPos] === s[colPos]) {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = Math.max(dp[i][j-1], dp[i-1][j]); 
            }
        }
    }
    
    return dp[s.length][s.length];
};