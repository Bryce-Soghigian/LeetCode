/**
 * @param {string} firstWord
 * @param {string} secondWord
 * @param {string} targetWord
 * @return {boolean}
 */
var isSumEqual = function(firstWord, secondWord, targetWord) {
    //
    const alpha={}
    const alphabet ="abcdefghijklmnopqrstuvwxyz"
    for(let i = 0;i<26;i++){
        alpha[alphabet[i]] = i
    }
    let firstWordStr = ""
    for(let i = 0;i<firstWord.length;i++){
        firstWordStr += alpha[firstWord[i]]
    }
    let sWordStr = ""
    for(let i = 0;i<secondWord.length;i++){
        sWordStr += alpha[secondWord[i]]
    }
    let target = ""
        for(let i = 0;i<targetWord.length;i++){
        target += alpha[targetWord[i]]
    }
    console.log(firstWordStr, sWordStr, target)
    return Number(target) === (Number(sWordStr) + Number(firstWordStr))
};