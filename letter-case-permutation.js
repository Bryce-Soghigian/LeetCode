/**
 * @param {string} S
 * @return {string[]}
 */
var letterCasePermutation = function(S) {
/**d

**/
const set = new Set();
const queue = [];
queue.unshift(S)
    
while(queue.length !== 0 ){
   let curr = queue.pop()
   console.log(curr)
   //Loop through curr and change each character to lower and to upper
  for(let i = 0;i<curr.length;i++){
      let uC = curr[i].toUpperCase(); // uC is the current char converted to uppercase
     let uL = curr[i].toLowerCase();
      let end = 0;
      if(i !== 0) end = i-1
      let upper = curr.slice(0,i) + uC + curr.slice(i+1,curr.length)
      let lower = curr.slice(0,i) + uL + curr.slice(i+1,curr.length)
      if(set.has(lower) === false){
          queue.unshift(lower)
          set.add(lower)
      }
      if(set.has(upper) === false){
          queue.unshift(upper)
          set.add(upper)
      }
        
  }
}
    

return Array.from(set)
};