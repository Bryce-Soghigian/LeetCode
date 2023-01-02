# class FirstUniqueTwoSets:
#     def __init__(self, nums: List[int]):
#         self.size = len(nums)
#         self.freq = {} # num : occurences, first_occ 
#         for i,num in enumerate(nums):
#             if num in self.freq: self.freq[num][0] += 1 
#             else: self.freq[num] = [1, i]
        
#         self.unique = {k for k, v in self.freq.items() if v[0] == 1}

#     def showFirstUnique(self) -> int:
#         """
#         O(len(self.unique))
#         """
#         if len(self.unique) == 0:
#             return -1
#         min_val = None
#         minimum = float('inf')
#         for val in self.unique:
#             if self.freq[val][1] < minimum:
#                 min_val = val
#                 minimum = self.freq[val][1]

#         return min_val
            

#     def add(self, value: int) -> None:
#         #O()
#         self.size += 1
#         if value in self.freq:
#             self.freq[value][0] += 1
#             if value in self.unique:
#                 self.unique.remove(value)
#         else:
#             self.freq[value] = [1, self.size-1]
#             self.unique.add(value)


class TreeNode:
    def __init__(self, val):
        self.val = val
        self.prev = None
        self.next = None
        
class FirstUnique:

    def __init__(self, nums: List[int]):
        self.head = TreeNode(0)
        self.tail = TreeNode(0)
        self.head.next, self.tail.prev = self.tail, self.head
        self.removed = set()
        self.seen = {}
        
        for num in nums:
            self.add(num)

    def showFirstUnique(self) -> int:
        if self.seen:
            return self.head.next.val
        return -1        

    def add(self, value: int) -> None:
        if value in self.removed: return
        if value not in self.seen:
            node = TreeNode(value)
            prev = self.tail.prev
            prev.next, self.tail.prev = node, node
            node.prev, node.next = prev, self.tail            
            self.seen[value] = node
        else:
            # delete node
            node = self.seen[value]
            prev, next = node.prev, node.next
            prev.next, next.prev = next, prev
            del self.seen[value]
            self.removed.add(node.val)