/**
 * @param {string} s
 * @return {string[]}
 */
var restoreIpAddresses = function(s) {
    if(s.length < 4 || s.length > 12) return [];
    var isValid = function(startIndex, endIndex) {
        if(endIndex - startIndex > 2 || (endIndex > startIndex && s[startIndex] === '0')
            || +s.slice(startIndex, endIndex + 1) > 255) {
            return false;
        }
        return true;
    }
    
    var result = [];
    var path = [];
    
    var backtrack = function(startIndex) {
        if(startIndex === s.length && path.length === 4) {
            result.push(path.join('.'));
        }
        for(var i = startIndex; path.length < 4 && s.length - i >= 4 - path.length; i++) {
            if(isValid(startIndex, i)) {
                path.push(s.slice(startIndex, i + 1))
                backtrack(i + 1);
                path.pop();
            } else {
                continue;
            }
        }
        return;
    }
    backtrack(0);
    return result;
};