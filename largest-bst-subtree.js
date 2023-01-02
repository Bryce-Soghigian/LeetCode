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
var largestBSTSubtree = function(root) {
       if(!root) {
        return 0;
    }
    let max = 0;
    let recurse = function(root) {
        if(!root) {
            return {size: 0, min: Infinity, max: -Infinity};
        }
   
        let left = recurse(root.left);
        let right = recurse(root.right);
        
        if(left.size === -1 || right.size === -1 || root.val <= left.max || root.val >= right.min) {
            return {size: -1, min: 0, max: 0};
        }
        let size = left.size + right.size + 1;
        max = Math.max(size, max);
        return {size: size, min: Math.min(left.min, root.val), max: Math.max(right.max, root.val)};
    }
    
    recurse(root);
    return max; 
};