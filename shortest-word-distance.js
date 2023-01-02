/**
 * @param {string[]} words
 * @param {string} word1
 * @param {string} word2
 * @return {number}
 */
var shortestDistance = function(words, word1, word2) {
    //Step 1 Find the ind of both words
    let wordOneInd = []
    let wordTwoInd = []
    for(let i = 0;i<words.length;i++){
        let curr = words[i];
        
        if(curr === word1){
            wordOneInd.push(i)
        }
        if(curr === word2){
            wordTwoInd.push(i)
        }
    }
    //We want to find the smallest difference between the two arrays
    let globalMin = Infinity;
    for(let i = 0;i<wordOneInd.length;i++){
        for(let j = 0;j<wordTwoInd.length;j++){
            let max = Math.max(wordOneInd[i],wordTwoInd[j])
            let min = Math.min(wordOneInd[i],wordTwoInd[j])
            let res = Math.abs(max-min)
            if(res < globalMin){
                globalMin = res
            }
        }
    }
    
    return globalMin
    
};