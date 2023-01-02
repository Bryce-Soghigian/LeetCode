/**
 * @param {number[][]} M
 * @return {number}
 */
var findCircleNum = function(M) {
    const dfs = (M,visited, i) => {
        for(let j = 0;j<M.length;j++){
if(M[i][j] === 1 && visited[j] == 0){
    visited[j] = 1
    dfs(M,visited,j)
}
        }
    }

       let visited = new Array(M.length).fill(0)
       let count = 0
       for(let i = 0;i<M.length;i++){
           if(visited[i] === 0){
               dfs(M,visited,i)
               count++
           }
       }
    return count
};