/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(n) {
 
    var reverseN = 0;
    var sign = n < 0;
    n = Math.abs(n);
    while (n) {
        reverseN = reverseN*10 + (n % 10);
        n = Math.floor(n/10);
    }
    return reverseN > 0x7FFFFFFF ? 0 : sign ? -reverseN : reverseN;



    /*const reversed = x.toString().split('').reverse().join(''); // turn a number into a string, then turn it into an array to reverse. 
  return Math.sign(x) * parseInt(reversed);  
    
    /*
    let newString1 = Number(x.toString().split('').reverse().join(''));
      if(newString1<0){
        let newString2 = '-'+newString1;
          return newString2;
    }else{
     return newString1;   
    }*/
   
};

