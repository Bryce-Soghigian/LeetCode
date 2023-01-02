/**
 * @param {string[]} words
 * @return {number}
 */
var maxProduct = function(words) {
        words.sort((a,b) => b.length-a.length);
    let res = 0;
    for (let i = 0; i < words.length-1; i++) {
        for (let j = i+1; j < words.length; j++) {
            if (!common(words[i],words[j])) 
                res = Math.max(res,words[i].length*words[j].length);
        }
    }
    return res;
    
    function common(a,b) { // a's length is greater than equal to b's
        let m = new Map();
        for (let letter of a) m.set(letter,1);
        for (let letter of b) 
            if (m.has(letter)) return true;
        return false;
    }
};