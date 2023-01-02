/**
 * @param {number} n
 * @return {number}
 */
var minSteps = function(n) {
   let answer = 0;
    let d = 2;
    while(n> 1){
         while(n%d === 0){
        answer += d;
        n /= d;
        
    }   
    d += 1    
    }
    return answer
};