/**
 * @param {number} n
 * @return {string}
 */
var generateTheString = function(n) {
    let stringType = null
    if(n%2===0){
        stringType = "even"
    }else{
        stringType = "odd"
    }
    if(stringType === "even"){
        let result = "a".repeat(n-1)+"b"
        return result
    }else{
      return "a".repeat(n)  
    }
};