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
 * @return {boolean}
 */
var isUnivalTree = function(root) {
    if (root === null) return true
    let num = root.val
    let isUnivalued = true
    const dfs = tree => {
        if(tree === null) return
        if(tree.val !== num ){
            isUnivalued = false
        }
        if(tree.left) dfs(tree.left)
        if(tree.right) dfs(tree.right)
    }
    dfs(root)
    return isUnivalued
};


