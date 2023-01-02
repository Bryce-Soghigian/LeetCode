class ListNode:
    def __init__(self, val, next=None):
        self.val = val
        self.next = next
    

class LinkedList:
    def __init__(self, head):
        self.head = head
    
    def __add__(self, other):
        if not self.head and not other:
            return None
        if not self.head:
            return other.head
        if not other:
            return self.head
        
        c_l1, c_l2 = self.head, other.head
        output = ListNode(val="*")
        output_pt = output
        carry = 0
        
        while c_l1 or c_l2 or carry:
        
        
            operand_one, operand_two = c_l1.val if c_l1 else 0, c_l2.val if c_l2 else 0
            summation = operand_one + operand_two + carry

            if summation < 10:
                output_pt.next = ListNode(val=summation)
                carry = 0

            else:
                ref = str(summation)
                output_pt.next = ListNode(ref[1])
                carry = int(ref[0])


            output_pt = output_pt.next
            c_l1 = (c_l1.next if c_l1 else None)
            c_l2 = (c_l2.next if c_l2 else None)
        
        return output.next
        
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:

        return LinkedList(l1) + LinkedList(l2)
