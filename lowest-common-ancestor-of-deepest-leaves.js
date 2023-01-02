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
 * @return {TreeNode}
 */
var lcaDeepestLeaves = function(root) {

    let maxDepth = 0;
    let lca = null;
    
        function dfs(node, depth) {
        if (node == null) return depth - 1;
        
        maxDepth = Math.max(maxDepth, depth);
        
        const leftMaxDepth = dfs(node.left, depth + 1);
        const rightMaxDepth = dfs(node.right, depth + 1);
        
        if (leftMaxDepth == maxDepth && rightMaxDepth == maxDepth) {
            lca = node;
        }
        
        return Math.max(leftMaxDepth, rightMaxDepth);
    }

    
    dfs(root, 0);
    
    return lca;
    

    
};