/**
 * @param {string} n
 * @param {number} x
 * @return {string}
 */
var maxValue = function(n, x) {
    let i = 0;
    if (n[0] !== '-') {
        for (; i < n.length; i++) {
            if (Number(n[i]) < x) break;
        }        
    } else {
        for (i = 1; i < n.length; i++) {
            if (Number(n[i]) > x) break;
        }
    }
    return n.slice(0, i) + x + n.slice(i)
};