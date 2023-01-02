/**
 * @param {number[][]} mat
 * @return {number[][]}
 */
var updateMatrix = function(matrix) {
    let i, l, j, m;
    
    let q = [];        
    
    // Find all zeroes in the matrix
    for (i = 0, l = matrix.length; i < l; i++) { 
        for ( j = 0, m = matrix[0].length; j < m; j++) {   
            if (matrix[i][j] === 0) {
                // Note the third param here, a zero to keep track of which "level" we're at. 
                // The zeroes are obviously at zero. 
                // Later in the bfs we will increase this for each unvisited neighbor
                q.push([i, j, 0]);
            } else {
                matrix[i][j] = Infinity;
            }  
        } 
    }
    
    // little helper array to find neighbors in a quick forEach loop.
    let dir = [[1,0],[0,1],[-1,0],[0,-1]];
    
    // Start BFS. BFS is the right choice so we minimize attempted double visits
    // BFS is like a stain that spreads, while DFS is like tendrils reaching out.
    while (q.length) {
        let pos = q.shift();
        
        // write value if we find it's lower than current (like those infinities)
        if (matrix[pos[0]][pos[1]] > pos[2]) {
            matrix[pos[0]][pos[1]] = pos[2];
        }
        
        // Look at all neighbor positions. Are they on the board? Are they not yet visited?
        // If yes to both, add to the q, with an increased "level" param at pos [2] 
        dir.forEach(function(d) {
            let next = [pos[0] + d[0], pos[1] + d[1], pos[2] + 1];
            // valid next coordinates?
            if (next[0] > -1 && next[0] < matrix.length && next[1] > -1 && next[1] < matrix[0].length) {
                // not yet marked?
                if (matrix[next[0]][next[1]] === Infinity) {
                    // add to q, but with increased index, which we stored at pos[2]
                    q.push(next);
                }
            }
        });
    }
    return matrix;
//     const bfs = (mat,i,j) => {
//         let seen = new Set();
//         const dir = [[-1,0],[1,0],[0,-1],[0,1]]
//         const queue = []
//         queue.push([i,j])
//         let closestZero = 0
//         while(queue.length !== 0){
//             //Search each level until we find a one and keep track of each level
//             let curr = queue.pop()
//             if(matrix[curr[0][curr[1]]] === 0){
//                 return closestZero
//             }else{
//                 closestZero++
//                 for(let c in dir){
//                     let x = curr[0] + dir[c][0]
//                     let y = curr[1] + dir[c][1]
//                     if(x < 0 || x >= mat.length || y < 0 || y >= mat[x].length || mat[x][y] === 0 || seen.has(`${x},${y}`)){
//                          continue
//                     }else{
//                         seen.add(`${x},${y}`)
//                         queue.unshift([x,y])
//                     }
//                 }
//             }
//         }
//         return closestZero
        
//     }
//     const matrix = mat
    
//     for(let i = 0;i<mat.length;i++){
//         for(let j = 0;j<mat[i].length;j++){
//             if(mat[i][j] === 1){
//                 let result = bfs(mat,i,j)
//                 matrix[i][j] = result
//             }else{
//               matrix[i][j] = 0
//             }
//         }
//     }
// return matrix
    
    
    
    
    
};