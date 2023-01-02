/**
 * @param {string} s
 * @return {string}
 */
var makeGood = function(s) {
    //While there are adj we want to remove them
    /**
    l
    e E
    if(same char and same letter casing){
    remove current
    restart loop
    }
    return s
     


    **/
    function checkCase(ch) {
    if(typeof(ch) !== "string"){
    return null
    }else {
      if (ch === ch.toUpperCase()) {
         return 'u';
      }
      if (ch == ch.toLowerCase()){
         return 'l';
      }
   }
}
    
let i = 1;
let stack = [];
stack.push(s[0])
while(i < s.length){

    let removed = stack.pop()

if(checkCase(removed) === "l" && checkCase(s[i]) === "u" && s[i].toLowerCase() === removed){

    s = s.slice(0,i-1) + s.slice(i+1,s.length)
    console.log("HIT") 
    i = 0
}else if (checkCase(removed) === "u" && checkCase(s[i]) === "l" && s[i].toUpperCase() === removed){

    s = s.slice(0,i-1) + s.slice(i+1,s.length)
    console.log("HIT") 
    i = 0;
}else{
    stack.push(s[i])
      i++
}
}
return s
};