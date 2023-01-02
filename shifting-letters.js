/**
 * @param {string} s
 * @param {number[]} shifts
 * @return {string}
 */
var shiftingLetters = function(S, shifts) {
      let alphabet = "abcdefghijklmnopqrstuvwxyz"
    var result = [];
    S.split("").forEach(i=>result.push(alphabet.indexOf(i)));
    let shiftsRepeat = shifts.reduce((a,b)=> a+b);
    let shiftsvalue = 0
    for(i=0;i<result.length;i++){
           result[i]+=shiftsRepeat
           shiftsRepeat-=shifts[shiftsvalue]
           shiftsvalue++;
           if(result[i]>26) result[i]%=26
           result[i]=alphabet[result[i]] 
        } 
     return result.join("")    
};