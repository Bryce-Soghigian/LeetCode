/**
 * @param {string} S
 * @param {number} K
 * @return {number}
 */
var numKLenSubstrNoRepeats = function(S, K) {
    if(K > S.length) return 0
let numOfValidStrings = 0
let start = 0
for(let end = K;end<=S.length;end++){
    let window = S.substring(start,end)
    let set = new Set(window.split(""))
    console.log(window,set)
    if(window.length === set.size){
        numOfValidStrings++
    }
    start++
}
    return numOfValidStrings
};