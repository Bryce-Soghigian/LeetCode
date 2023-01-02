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
 * @return {number}
 */
var findSecondMinimumValue = function(root) {
    //Find all of the values
    
    const seen = new Set()
    
    
    const dfs = tree => {
        if(tree === null) return 
        seen.add(tree.val)
        
        if(tree.left) dfs(tree.left)
        if(tree.right) dfs(tree.right)
    }
    dfs(root)
    let array = Array.from(seen).sort((a,b) => a-b)
    return array.length === 1 ? -1 : array[1]
};