class Solution:
    def findOriginalArray(self, changed: List[int]) -> List[int]:
        changed.sort()
        odd, even = collections.defaultdict(int), collections.defaultdict(int)
        res = []
        for num in changed:
            if num % 2:
                odd[num] += 1
            else:
                if num // 2 in odd and odd[num // 2] != 0:
                    odd[num // 2] -= 1
                    res.append(num//2)
                elif num // 2 in even and even[num // 2] != 0:
                    even[num // 2] -= 1
                    res.append(num//2)
                else:
                    even[num] += 1
        if (odd and any(odd.values())) or (even and any(even.values())):
            return []
        return res
                       