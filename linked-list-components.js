/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number[]} nums
 * @return {number}
 */
var numComponents = function(head, G) {
        let count = 0;
        let node = head;
        let set = new Set()
        for (let i in G) {
            set.add(G[i]);
        }
    
        while (node != null) {
            if (set.has(node.val)) {
                count++;
                while (node !== null && set.has(node.val)) {
                    node = node.next;
                }
            } else {
                node = node.next;
            }
        }
        return count;
}