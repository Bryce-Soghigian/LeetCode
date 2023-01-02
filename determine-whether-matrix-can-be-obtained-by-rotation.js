/**
 * @param {number[][]} mat
 * @param {number[][]} target
 * @return {boolean}
 */
var findRotation = function(mat, target) {
    //We rotate the mtrix 90 degrees 4 times and check if it is equal
    
var rotate = function(matrix) {
//o(n * n)  

let length = matrix.length-1;
for(let i =0; i<length;i++ ){
    for(let j = i; j<length-i;j++){
        let temp = matrix[i][j]
        matrix[i][j] = matrix[length - j][i]
        matrix[length-j][i] = matrix[length-i][length-j]
        matrix[length - i ][length - j] = matrix[j][length-i]
        matrix[j][length - i] = temp
    }
}
    return matrix
};
let i = 0 
let rotated = mat
while(i !== 4){
    //Compare all values in matrix
    let isM = true
    for(let i = 0;i<rotated.length;i++){
        for(let j = 0;j<rotated[i].length;j++){
            if(rotated[i][j] !== target[i][j] ){
                isM = false
                break
            }
        }
    }
    if(isM === true){
        return true
    }else{
        rotated = rotate(rotated)
        i++
    }
    
}
return false
    
    
};