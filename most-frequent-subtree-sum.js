/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var findFrequentTreeSum = function(root) {
    
    const getSum = (tree, sum = []) => {
        if(tree){
            sum.push(tree.val)
            getSum(tree.left,sum)
            getSum(tree.right,sum)
        }
        return sum
    }
    const add = (array ) => {
        let sum = 0
        for(let i = 0;i<array.length;i++){
            sum += array[i]
        }
        return sum
    }
    let hash = {}
    const traverse = (tree) => {
        if(tree === null) return
        let total = add(getSum(tree))
        if(total in hash){
            hash[total] = hash[total] + 1
        }else{
            hash[total] = 1
        }
        if(tree.left) traverse(tree.left)
        if(tree.right) traverse(tree.right)
    }
    traverse(root)
    //We want to grab the most frequent sum
    let max = Math.max(...Object.values(hash))
    let results = []
    for(let key in hash){
        if(hash[key] === max){
            results.push(key)
        }
    }
    return results

}