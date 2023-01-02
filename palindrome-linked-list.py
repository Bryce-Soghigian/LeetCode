class ListNode:
    def __init__(self, val, next=None):
        self.val = val
        self.next = next
class LinkedList:
    def __init__(self, head):
        self.head = head
    def reverse(self):
        prev = None
        curr = self.head
        while curr:
            temp = curr.next
            curr.next = prev
            prev = curr
            curr = temp
        self.head = prev
    
    @property
    def val(self):
        if not self.head:
            return None
        return self.head.val
    
    def _next(self):
        if not self.head:
            return
        self.head = self.head.next
        return
            
class Solution:
    def find_middle_node(self, head):
        fast = head
        slow = head
        while fast.next is not None and fast.next.next is not None:
            fast = fast.next.next
            slow = slow.next
        return slow
    
    def isPalindrome(self, head):
        
        reversed_half = LinkedList(self.find_middle_node(head))
        reversed_half.reverse()
        
        
        while reversed_half.val is not None and head:
            if reversed_half.head is None:
                return True
            if reversed_half.val != head.val:
                return False
            
            reversed_half._next()
            head = head.next
        return True