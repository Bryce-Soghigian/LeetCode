var mostCompetitive = function(N, K) {
    let len = N.length, ans = [N[0]], curr = 0, rems = len - K
    for (let i = 1; i < len;)
        if (!rems) return ans.concat(N.slice(i))
        else if (N[i] < ans[curr]) ans.pop(), curr--, rems--
        else ans.push(N[i]), i++, curr++
    return ans.slice(0,K)
};