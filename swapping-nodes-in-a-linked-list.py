class Solution:
    def swapNodes(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        cur_node = head
        list_len = 0
        while cur_node:
            list_len += 1
            if list_len == k:
                front_node = cur_node
                end_node = head
            if list_len > k:
                end_node = end_node.next
            cur_node = cur_node.next
        front_node.val, end_node.val = end_node.val, front_node.val
        return head