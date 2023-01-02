/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {

// let str = s.split(" ")
// let index = str.length
// if(str.length ===" "||str.length === ""){
//     if(str[index-2].length){
//       return str[index-2].length  
//     }
    
    
// }else{
   
//     return str[index-1].length
// }
    var wordArray = s.split(" ")
var lastWord = null;

for(var i = wordArray.length-1;i>=0;i--){
if(wordArray[i] !== ""){
lastWord = wordArray[i];
break;
}
}

return lastWord? lastWord.length:0;
    


};