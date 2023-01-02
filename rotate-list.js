/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var rotateRight = function(head, k) {
    if (!head) return head
    const length = lengthOfList(head)
    for (let i = 0; i < (k % length); i++) {
        let newHead = new ListNode(), node = head, prev = head
        while (node.next) {
            prev = node
            node = node.next
        }
        if (length > 0) {
            prev.next = null
            newHead.val = node.val
            newHead.next = head
            head = newHead
        }
    }
    return head
};

const lengthOfList = head => {
    let node = head, length = 0
    while (node) {
        node = node.next
        length++
    }
    return length
}

// /**
//  * Definition for singly-linked list.
//  * function ListNode(val, next) {
//  *     this.val = (val===undefined ? 0 : val)
//  *     this.next = (next===undefined ? null : next)
//  * }
//  */
// /**
//  * @param {ListNode} head
//  * @param {number} k
//  * @return {ListNode}
//  */
// var rotateRight = function(head, k) {
//     if(head === null) return null
//     let newList = []
//     let curr = head;
//     while(curr){
//         newList.push(curr.val)
//         curr = curr.next
//     }
//     //Then we add the last k nodes to the front of our new list
//     //then make a list of all the remaining nodes
    
//     while(k !==0){
//         k--
//         let removed = newList.pop()
//         newList.unshift(removed)
//     }
//     let newHead = new ListNode(newList[0])
//     let pointer = newHead;
//     for(let i = 1;i<newList.length;i++){
//         let newNode = new ListNode(newList[i])
//         pointer.next = newNode
//         pointer = pointer.next
//     }

//     return newHead
// };