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
 * @param {number} target
 * @return {number}
 */
var closestValue = function(root, target) {
    //Traverse the whole tree
    //For each node check if math.abs(tree.val - target)
    let closest = Infinity
    let minDiff = Infinity
    const traverse = (tree) => {
        if(tree === null) return
        let diff = Math.abs(tree.val - target)
        console.log(diff,tree.val)
        if(diff <= minDiff){
            closest = tree.val
            minDiff = diff
        }
        if(tree.left) traverse(tree.left)
        if(tree.right) traverse(tree.right)
        
    }
    traverse(root)
    return closest
};