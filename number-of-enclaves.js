/**
 * @param {number[][]} A
 * @return {number}
 */
var numEnclaves = function(matrix) {

    //console.log(matrix);
    
    let visited = [];

    for(let i = 0 ; i <= matrix.length-1 ; i++)
    {
        visited.push(new Array(matrix[0].length-1).fill(false));
    }

    for(let i = 0 ; i<= matrix.length-1 ; i++)
    {
        for(let j = 0 ; j<= matrix[0].length-1 ; j++)
        {
            if( (i == 0 || i == matrix.length-1 || j == matrix[0].length-1 || j == 0) && matrix[i][j] ==1)
            {
                cleanMatrix( i,j,matrix, visited);
            }
        }
    }
    
    //console.log(matrix);
    let counter = 0;

    for(let i = 0 ; i<= matrix.length-1 ; i++)
    {
        for(let j = 0 ; j<= matrix[0].length-1 ; j++)
        {
            if(matrix[i][j] == 1)
            {
                counter++;
            }
        }
    }

    return counter;


};



function cleanMatrix( i, j, matrix, visited)
{
    if(i < 0 || i > matrix.length-1 || j < 0 || j > matrix[0].length-1 || visited[i][j] || matrix[i][j] == 0 || matrix[i][j] == -1)
    {
        return;
    }

    visited[i][j] = true;
    matrix[i][j] = -1;

    cleanMatrix( i+1, j, matrix, visited)
    cleanMatrix( i-1, j, matrix, visited);
    cleanMatrix( i, j+1, matrix, visited);
    cleanMatrix( i, j-1, matrix, visited);
    
}