/**
 * @param {number} x
 * @param {number} y
 * @param {number} bound
 * @return {number[]}
 */
var powerfulIntegers = function(x, y, bound) {
/*

Loop through the range of its an find powerful numbers

a powerful number is x^i + y^j

take x n y and raise them to the possible powers

if powerful num is less than bound
push powerful number to output



*/
let output = new Set()
    const xLimit = Math.log(bound) / Math.log(x);
    const yLimit = Math.log(bound) / Math.log(y);

for(let i = 0;i<xLimit;i++){
    for(let j = 0;j<yLimit;j++){
    let num = Math.pow(x,i) + Math.pow(y,j)
    if(num <= bound){
        output.add(num)
    }
    if(y === 1) break;

    }
    if(x === 1) break
}


return Array.from(output)
};