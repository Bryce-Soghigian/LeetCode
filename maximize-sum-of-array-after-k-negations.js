/**
 * @param {number[]} A
 * @param {number} K
 * @return {number}
 */
var largestSumAfterKNegations = function(A, K) {
    const pos = [];
    const neg = [];
    
    for(let num of A) {
        if (num < 0) {
            neg.push(num);
        } else {
            pos.push(num);
        }    
    }
    neg.sort(sortNum);
    
    while(K > 0 && neg.length > 0) {        
        pos.push(-neg.pop());
        K--;
    }
    
    pos.sort(sortNum);
    if (K % 2 == 1) {
        pos[pos.length - 1] *= -1;
    }
    
    return pos.concat(neg).reduce((sum, num) => sum += num, 0);
};

const sortNum = function(a, b) {
    return b - a;
}