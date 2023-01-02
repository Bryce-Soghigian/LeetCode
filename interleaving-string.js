const initialize2DArrayNew = (n, m) => { let data = []; for (let i = 0; i < n; i++) { let tmp = Array(m).fill(false); data.push(tmp); } return data; };

const isInterleave = (s1, s2, s3) => {
    let [n1, n2, n3] = [s1.length, s2.length, s3.length];
    if (n1 + n2 != n3) return false;
	let dp = initialize2DArrayNew(n1 + 1, n2 + 1);
    for (let i = 0; i <= n1; i++) {
        dp[i][0] = s1.slice(0, i) == s3.slice(0, i);
    } 
    for (let i = 0; i <= n2; i++) {
        dp[0][i] = s2.slice(0, i) == s3.slice(0, i);
    }
    for (let i = 1; i <= n1; i++) {
        for (let j = 1; j <= n2; j++) {
            dp[i][j] = (dp[i - 1][j] && s1[i - 1] == s3[i + j - 1]) ||
                  (dp[i][j - 1] && s2[j - 1] == s3[i + j - 1]);
        }
    }
	return dp[n1][n2];
};