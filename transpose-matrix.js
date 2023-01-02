/**
 * @param {number[][]} matrix
 * @return {number[][]}
 */
var transpose = function(A) {
    if(A.length === 0) return []
        let res = [...Array(A[0].length)].map(() => Array(A.length).fill(0));

    for(let i=0;i<A.length;i++){
        for(let j=0;j<A[0].length;j++){            
            res[j][i]=A[i][j];
        }
    }
    return res;
};