/**
 * @param {number[][]} triangle
 * @return {number}
 */
var minimumTotal = function(triangle) {
   const N = triangle.length;
   const resultArray = Array.from({length: N+1}).map((item ,i) => Array.from({length: i+1}).fill(0));
   let minResult = Number.MAX_VALUE;
    
   for(let i = 1 ; i <= N ; i++){
       for(let j = 1; j <= i ; j++){
           if(i === 1){
               //starting point
               resultArray[1][1] = triangle[0][0];
           }else if(j===1){
               //triangle left side value
               resultArray[i][j] = resultArray[i-1][j] + triangle[i -1][j -1];
           }else if(i === j){
               //triangle right side value
               resultArray[i][j] = resultArray[i-1][j-1] + triangle[i -1][j -1];
           }else{
               //regular
              resultArray[i][j] = Math.min(resultArray[i-1][j-1] , resultArray[i-1][j]) + triangle[i -1][j -1];
           }
           
           if(i === N){
               minResult = Math.min(minResult , resultArray[i][j]);
           }
       }
   }
    
   return minResult;
    };

