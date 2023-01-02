class SLLNode:
    def __init__(self, val):
        self.val = val
        self.next = None
        
class Solution:
    def findTheWinner(self, n: int, k: int) -> int:
        queue = deque()
        
        for i in range(1, n+1):
            queue.append(i)
        
        while len(queue) > 1:
            removal = k - 1
            while removal > 0:
                queue.append(queue.popleft())
                removal -= 1
            queue.popleft()
        return queue.popleft()
#         head = SLLNode(val="sentinel")
#         curr = head
#         #1. build list
#         for i in range(1, n+1):
#             new_node = SLLNode(val=i)
#             curr.next = new_node
#             curr = curr.next
        
#         curr = head.next
        
#         while curr.next:
#             print(curr.val)
#             for i in range(k):
#                 curr = curr.next
#             if curr.next.next:
#                 curr.next = curr.next.next
#             if curr.val == "sentinel":
#                 curr = head.next
        
#         #2 play game
#             # delete the node
#         return head.next.val
    