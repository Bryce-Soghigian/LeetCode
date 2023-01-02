/**
 * @param {string} address
 * @return {string}
 */
var defangIPaddr = function(address) {
    
/**
We want to split the address into an array 
then we want to take that array and check for the "."
character
if(currentChar === "."){

}

    
    
    
**/
    let addressArray = address.split("");
    let defangedString = "";
    for(let i = 0;i<addressArray.length;i++){
        if(addressArray[i]==="."){
            defangedString += "[.]"
        }else{
            defangedString += addressArray[i]
        }
    }
    return defangedString
};