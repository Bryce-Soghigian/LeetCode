/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {ListNode} head
 * @return {TreeNode}
 */
var sortedListToBST = function(head) {
    let curr = head;
    let values = []
    while(curr){
        values.push(curr.val)
        curr = curr.next
    }

    
   function bst(nums,left,right){
        if(left > right){
            return null;
        }
        let mid = Math.floor(left + (right - left ) /2);
        let current = new TreeNode(nums[mid]);
        current.left = bst(nums,left,mid - 1);
        current.right = bst(nums,mid+1,right);
        
        return current;
    }

    return bst(values,0,values.length-1);

};