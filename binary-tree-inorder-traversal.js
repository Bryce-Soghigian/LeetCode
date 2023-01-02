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
var inorderTraversal = function(root) {
//  let result = []
//  let stack = []  
//  stack.push(root)
// while(stack.length !== 0){
//     //Check the top item of the stack for a left element
//     let curr = stack.pop()
//     if(curr.left){
//         stack.push(curr.left)
//     }
//     stack.pop()
//     result.push(curr.val)
//     if(curr.right){
//         stack.push(curr.right)
//     }
    
// }
    
//     return result
     let values = []
    
    
    const traverse = node => {
        if(node === null) return []
        if(node !== null){
            traverse(node.left)
            values.push(node.val)
            traverse(node.right)
        }

    }
    traverse(root)
    return values
};