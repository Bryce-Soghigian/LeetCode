/**
 * @param {number[]} arr
 * @param {number} k
 * @return {number}
 */
var findLeastNumOfUniqueInts = function(arr, k) {
    // o(n) + nlogn +o(k) + freq.length
    let freq = {}
    for(let i = 0;i<arr.length;i++){
        if(arr[i] in freq){
            freq[arr[i]] = freq[arr[i]] + 1
        }else{
            freq[arr[i]] = 1
        }
    }
    let MinHeap = Object.entries(freq)
    MinHeap.sort((a,b) => {
        if(a[1] > b[1]){
            return 1
        }else if(a[1] < b[1]){
            return -1
        }else if(a[1] === b[1]){
            return 0
        }
    })
    let i = 0;
    while(i < MinHeap.length && k !== 0){
        if(MinHeap[i][1]=== 0){
            i++
        }else if(MinHeap[i][1] > 0){
            k--
            MinHeap[i][1] = MinHeap[i][1] - 1
        }
    }
    let count = 0;
    for(let i = 0;i<MinHeap.length;i++){
        if(MinHeap[i][1] !== 0){
            count++
        }
    }
    return count
};