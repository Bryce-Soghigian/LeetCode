/**
 * @param {string} S
 * @return {number[]}
 */
var partitionLabels = function(S) {
   /**
   {a:[start,end]}
   1. Index all of our values
   2. Greedy algorithm that breaks the partition into as many parts as possible
   if the freq of a letter is greater than 2 from its start index to its end index it cannot be broken
   
   **/
    
    let output = []
    let map = {}
    for(let i = 0;i<S.length;i++){
        map[S[i]] = i 
    }
    
    let i = 0;
    let maxSize = 0;
    let last = -1;
    for(let i = 0;i<S.length;i++){
        if(maxSize < map[S[i]]){
            maxSize = map[S[i]]
        }
        if(maxSize === i ){
            output.push(i- last)
            last = i
        } 
    }
    
    return output
};