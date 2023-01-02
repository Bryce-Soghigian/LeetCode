class Solution:
    def findMid(self, node):
        slow = fast = node
        
        while(fast.next and fast.next.next):
            fast = fast.next.next
            slow = slow.next
        
        return slow
    
    def merge(self, node1, node2):
        dummy = cur = ListNode()
        intMax = float('inf')
        
        while(node1 or node2):
            value1, value2 = node1.val if(node1) else intMax, node2.val if(node2) else intMax
            if(value1 < value2):
                cur.next = node1
                node1 = node1.next
            else:
                cur.next = node2
                node2 = node2.next
            cur = cur.next
            
        return dummy.next
            
    def mergeSort(self, node):
		# Incase of single node or empty node we don't have to do anything.
        if(node is None or node.next is None):  return node
        
        mid = self.findMid(node) #Find the mid
        nextNode = mid.next #Get the start of second half 
        mid.next = None #Break at mid
        
        firstHalf = self.mergeSort(node)
        secondHalf = self.mergeSort(nextNode)
        
        return self.merge(firstHalf, secondHalf)
        
    
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        return self.mergeSort(head)
        
# class Solution:
#     def merge_two_sorted_lists(self, list1, list2):
        
#         output = ListNode("*")
#         output_pt = output
        
#         while list1 or list2:
#             if not list2:
#                 output_pt.next = ListNode(list1.val)
#                 list1 = list1.next
#                 output_pt = output_pt.next
#                 continue
#             if not list1:
#                 output_pt.next = ListNode(list2.val)
#                 list2 = list2.next
#                 output_pt = output_pt.next
#                 continue
            
#             if list1.val > list2.val:
#                 output_pt.next = ListNode(list2.val)
#                 list2 = list2.next
#             else:
#                 output_pt.next = ListNode(list1.val)
#                 list1 = list1.next
#             output_pt = output_pt.next
        
#         return output.next
                
#     def return_pairs_of_list(self, list_of_nodes):

#         output = deque()
#         if len(list_of_nodes) % 2 != 0:
#             output.appendleft([list_of_nodes.pop(),ListNode("ODD")])
#         for i in range(0,len(list_of_nodes), 2):
#             output.append((list_of_nodes[i], list_of_nodes[i+1]))
#         return output
            
#     def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
#         list_nodes = deque()
#         while head:
#             list_nodes.append(ListNode(val=head.val))
#             head = head.next
#         while len(list_nodes) > 1:
#             pairs = self.return_pairs_of_list(list_nodes)
#             list_nodes = deque()
#             if pairs[0][1].val == "ODD":
#                 list_nodes.appendleft(pairs[0][0])
#                 pairs.popleft()
#             for linked_list_a, linked_list_b in pairs:
#                 list_nodes.append(self.merge_two_sorted_lists(linked_list_a, linked_list_b))

#         return 