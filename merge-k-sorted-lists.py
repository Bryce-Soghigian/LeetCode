# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

from heapq import heappop, heappush, heapify
class Solution:
    def get_list(self, linked_list):
        
        while linked_list:
            yield linked_list.val
            linked_list = linked_list.next
    def build_list(self, nums):
        
        output = ListNode("*")
        pt = output
        for num in nums:
            pt.next = ListNode(val=num)
            pt = pt.next
            
        return output.next
        
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        output = []
        for l in lists:
            items = self.get_list(l)
        
            for item in items:
                output.append(item)
        output.sort()
        return self.build_list(output)

    
        
        