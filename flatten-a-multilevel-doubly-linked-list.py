# """
# # Definition for a Node.
# class Node:
#     def __init__(self, val, prev, next, child):
#         self.val = val
#         self.prev = prev
#         self.next = next
#         self.child = child
# """

# class Solution:
#     def flatten(self, head: 'Optional[Node]') -> 'Optional[Node]':
#         """
#         output = 1,2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6
#         [(3->4->5->6),]
#             [7 -> ]
#         1 2 3
#         append curr pointer to stack
#         8 -> 9 -> 10 
#         add a list to the stack
#         1,2,3,7,8
        
#         stack []
#         [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
        
#         we pop of stack when we are at a none value
#         we add onto stack when we are at a child value
#         3
#         """
        
#         output_head = Node(val="*")
#         ref = output_head
#         stack = deque()
#         stack.append(head)
#         while stack:
#             curr = stack.pop()
#             print(curr.val,"HEAD")
#             while curr:
#                 print(curr.val)
#                 if curr.child == None:
#                     new_node = Node(val=curr.val)
#                     output_head.next = new_node
#                     output_head = output_head.next
#                     curr = curr.next
                    
#                 else:
#                     print("child")
#                     output_head.next = Node(val =curr.val)
#                     output_head = output_head.next
#                     stack.append(curr.child)
#                     break
                    
#         return ref.next
class Solution:
    def flatten(self, head: 'Optional[Node]') -> 'Optional[Node]':
        tmp = head
        
        ##iterate over the linkedlist
        while(tmp):
            ##store next and previous pointer in case we encounter the child
            nxt = tmp.next
            prv = tmp.prev
            ##if the current node has child
            if(tmp.child):
                node = tmp.child
                tmp.next = tmp.child
                node.prev = tmp
                ##iterate till an end of the child (layer-2 linked list)
                while(node.next != None):
                    node = node.next
                
                ##link the last node in layer-2 to the next of the node of which it is child 
                node.next = nxt
                if nxt !=  None:
                    nxt.prev = node
                tmp.child = None
            
            tmp = tmp.next

        return head