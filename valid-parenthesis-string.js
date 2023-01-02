/**
 * @param {string} s
 * @return {boolean}
 */
var checkValidString = function(s) {
    if (!s) return true;
    var stack = [],
    l=0,r=0,len = s.length,star = 0;
    for (let i = 0; i < len; i++) {
        if (s[i] === "(") {
            l++;
            stack.push(true);
        } else if (s[i] === ")") {
            r++;
            if(star + l < r){
                return false;
            }
            stack.pop();
        } else if (s[i] === "*") {
            star++;
            stack.pop();
        }
    }
    if (stack.length === 0) return true;
    return false;
};