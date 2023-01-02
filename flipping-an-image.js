/**
 * @param {number[][]} A
 * @return {number[][]}
 */
var flipAndInvertImage = function(A) {
    //=======Flipping It===============//
for(let i = 0; i<A.length;i++){
        A[i].reverse();

};
//     let new_arr = A.map(x=> {
//         x.reverse()
//     }).map(z => {
//         let temp = []

            
//             if(z === 1){
//             temp.push(0)
//             }
//             if(z===0){
//                temp.push(1)
//             }
            

//         return temp
//     })

// return new_arr
    //==========Reversing the Characters==========//
    let new_arr = []
    for(let i = 0;i<A.length;i++){

        let temp_arr = []
        for(let j = 0;j<A.length;j++){
            
            if(A[i][j]===1){
                
                temp_arr.push(0)
            
            }
            if(A[i][j]===0){
               temp_arr.push(1)
            }
            
        }
        new_arr.push(temp_arr)
    }
    console.log(A,"current")
    //
    console.log(new_arr,"new_arr")
    return new_arr
}