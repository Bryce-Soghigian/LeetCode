/**
 * @param {number[]} costs
 * @param {number} coins
 * @return {number}
 */
var maxIceCream = function(costs, coins) {
    /*
    */
    let ans = 0
    let tot = 0
    costs = costs.sort((a,b) => a-b)
    for(let i = 0;i<costs.length;i++){
        tot += costs[i]
        if(tot>coins){
            break
        }
        ans++
    }
    return ans
};