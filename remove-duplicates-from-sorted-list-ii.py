class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        
        prev = None
        node = head
        while node and node.next:
            
            skipped = node.next
            while skipped and node.val == skipped.val:
                skipped = skipped.next
                
            if node.next == skipped:
                prev = node
                node = node.next
            else:
                if prev == None:
                    head = node = skipped
                else:
                    prev.next = skipped
                    node = skipped

        
        return head