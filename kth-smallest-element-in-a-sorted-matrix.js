/**
 * @param {number[][]} matrix
 * @param {number} k
 * @return {number}
 */
var kthSmallest = function(matrix, k) {
let heap = []
for(let i = 0;i<matrix.length;i++){
    for(let j = 0;j<matrix[i].length;j++){
        const current = matrix[i][j]
        heap.push(current)
    }
}
heap.sort((a,b) => a-b) 


return heap[k-1]
    
};