/**
 * @param {string} secret
 * @param {string} guess
 * @return {string}
 */
var getHint = function(secret, guess) {
  let secretLength = secret.length;
    let hash = {};
    let cows = 0;
    let bulls = 0
    //Loop through secret an count the secret through our hash
    // 
    for(let i = 0;i<secretLength;i++){
        let item = secret[i]
        if(hash[item]){
            hash[item] += 1
        }else{
            hash[item] = 1
        }
    }
    for(let i =0;i<secretLength;i++){
        //Calculate the guess
        let g = guess[i]
        if(hash[g]){
            cows += 1
            hash[g] -= 1
        }
        if(secret[i] === guess[i]){
            bulls += 1
        }
    }
    return `${bulls}A${cows-bulls}B`
};