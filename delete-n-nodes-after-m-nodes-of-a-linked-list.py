
class Solution(object):
    def deleteNodes(self, head, m, n):
        """
        :type head: ListNode
        :type m: int
        :type n: int
        :rtype: ListNode
        """
        curr = head
        while curr:
            count = 1
			# skip m nodes
            while(curr and count<m):
                curr = curr.next
                count+=1
            count = 0
			#delete n nodes
            while(curr and curr.next and count<n):
                curr.next = curr.next.next
                count+=1
            if curr:
                curr = curr.next
        return head