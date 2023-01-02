/**
 * @param {string} s
 * @return {string}
 */
var removeVowels = function(s) {
    let vowels = "aeiou";
    let dict = new Set()
    for(let i =0;i<vowels.length;i++){
        dict.add(vowels[i])
    }
    
    let newStr = ''
    s.split("").map(x => {
        if(dict.has(x)===false){
            newStr += x
        }
    })
    return newStr
};