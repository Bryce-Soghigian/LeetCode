/**
 * @param {string[]} words
 * @param {number} maxWidth
 * @return {string[]}
 */
var fullJustify = function(words, maxWidth) {
        let cur = [], charWidth = 0;
    let result = [];
    for (let word of words) {
        if (charWidth + word.length + cur.length > maxWidth) { // exceeding the maxWidth 
            let space = maxWidth - charWidth;
            let padded = [];
            if (cur.length == 1) {
                padded = [cur[0]+=' '.repeat(space)];
            } else {
                var spread = parseInt(space/(cur.length-1)); // shared space counts
                var extra = space % (cur.length - 1); // extra soace counts
                for (let j=0;j<cur.length;j++) {
                    let rep = 0;
                    if (j!=cur.length-1) rep = spread + (extra-->0?1:0)
                    padded.push(cur[j]+=" ".repeat(rep));
                }
            }
            result.push(padded.join(''));
            cur = [];
            charWidth = 0;
        } 
        cur.push(word);
        charWidth+=word.length;
    }
    result.push(cur.join(' ').padEnd(maxWidth, ' '));
    return result;
};