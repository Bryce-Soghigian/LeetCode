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
var convertBST = function(root) {
        let total = 0
        const traverse = (tree ) => {
            if(tree === null) return 
            if(tree.right) traverse(tree.right )
                    total+= tree.val
                    tree.val = total 
            if(tree.left) traverse(tree.left)
        }
        traverse(root,0)
        return root
};