/**
 * @param {number[]} arr
 * @return {number}
 */
var countElements = function(arr) {
    /*
    put all the values in a hashtable
    if x + 1 exists in the hashtable 
    increment final count
    walk through the array
    for each element count 
    
    */
    let hash = new Set(arr)
    let count = 0
    arr.map(ele => {
        if(hash.has(ele+1)){
            count++
        }
    })
    return count
};