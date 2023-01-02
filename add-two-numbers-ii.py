# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def _reverse_list(self, linked_list,):
        """
        3 -> 2 -> 1
        
          c
      p  
        Move forward curr pointer
        temp = curr
        set prev = curr
        set prev.next = temp
        """
        prev = None
        curr = linked_list
        
        while curr:

            temp = curr.next # None
            
            curr.next = prev # 4
            prev = curr # 5
            curr = temp # None
        return prev
            

        return prev
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        """
        Input: l1 = [7,2,4,3], 
                 l2 = [5,6,4]
                     s
                     
        Approaches
        1. Reverse the two linked lists and suport elementary Math for the linked list
        
        3 -> 4 -> 2 -> 7
        4 -> 6 -> 5
        
        
        output -> 7 -> 0 -> 8 -> 7
        operand_one = 7
        operand_two = None == 0
        carry = 0
        
        1. Reverse all lists
        2. Generate two pointers
        3. Perform elementary math on the two linked lists
        
        
Output: [7,8,0,7]
        
        """
        output = ListNode("*")
        
        l1 = self._reverse_list(l1)
        l2 = self._reverse_list(l2)
        operand_one = l1
        operand_two = l2
        output_pt = output
        carry = 0
        while operand_one or operand_two:
            operand_one_val = operand_one.val if operand_one is not None else 0
            operand_two_val = operand_two.val if operand_two is not None else 0
            total = operand_one_val + operand_two_val + carry
            if total - 9 > 0:
                carry_parts = str(total)
                carry = int(carry_parts[0])
                output_pt.next = ListNode(val=int(carry_parts[1]))
                output_pt = output_pt.next
            else:
                output_pt.next = ListNode(val=total)
                output_pt = output_pt.next
                carry = 0
            
            
            if operand_one:
                operand_one = operand_one.next
            if operand_two:
                operand_two = operand_two.next
            
        if carry != 0:
            print(carry)
            output_pt.next = ListNode(val=carry)
        return self._reverse_list((output.next))