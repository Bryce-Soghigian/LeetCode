/**
 * @param {number} num
 * @return {number}
 */
var numberOfSteps  = function(num) {
    /**
    Input: num = 14
Output: 6
Explanation: 
Step 1) 14 is even; divide by 2 and obtain 7. 
Step 2) 7 is odd; subtract 1 and obtain 6.
Step 3) 6 is even; divide by 2 and obtain 3. 
Step 4) 3 is odd; subtract 1 and obtain 2. 
Step 5) 2 is even; divide by 2 and obtain 1. 
Step 6) 1 is odd; subtract 1 and obtain 0.

if(num is even){
//increment count
//We divide num by 2 and recurse

}
if(num is odd){
icrement count
//subtract 1
//recurse
}
    
    
    
    **/
 let count = 0;   
const reduceNumber = (num) => {
    if(num === 0){
        return
    }
    if(num %2 === 0){
       count++ 
       reduceNumber(num/2)
    }else{
        count++
        reduceNumber(num-1)
    }
}
reduceNumber(num)
    return count
};