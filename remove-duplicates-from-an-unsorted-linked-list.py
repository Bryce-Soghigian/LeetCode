class Solution:
    def deleteDuplicatesUnsorted(self, head: ListNode) -> ListNode:
        # edge case
        if head.next is None:
            return head

        # use ordered dictionary to check duplicate nodes
        d, dup_vals = collections.OrderedDict(), set()
        h = head
        while h:
            d[h.val] = d.get(h.val, 0)+1
            if d[h.val] > 1:
                dup_vals.add(h.val)
            h = h.next

        all_uniq_vals = list(d.keys())
        res, h = None, None
        count = 0
        for e in all_uniq_vals:
            if e not in dup_vals:
                if count == 0:
                    res = ListNode(val=e)
                    h = res
                else:
                    new_node = ListNode(val=e)
                    h.next = new_node
                    h = new_node
                count += 1
        return res