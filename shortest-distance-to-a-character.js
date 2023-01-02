/**
 * @param {string} s
 * @param {character} c
 * @return {number[]}
 */
var shortestToChar = function(s, c) {
    /**
    The first solution that comes to mind is that you could do a double for loop and search for the char and mark
    **/
    
    let shortestArray = []
    for(let i = 0;i<s.length;i++){
        let found = false
        if(s[i]===c){
            shortestArray.push(0)
        }else{
            
            for(let j = i;j<s.length;j++){
                if(s[j] === c){
                   shortestArray.push(Math.abs(j-i)) 
                    found = true
                    break
                }
            }
                    if(found === false) shortestArray.push(Infinity)

        }
    }
    
    for(let i = s.length;i >= 0;i--){
        if(shortestArray[i] !== 0){
            for(let j = i; j>= 0;j--){
                if(s[j] === c ){
                    if(Math.abs(j-i) < shortestArray[i])
                    shortestArray[i] = Math.abs(j-i)
                    break
        
                }
            }
        }
        
        
    }
    return shortestArray
};