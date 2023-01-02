class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        jewels = {jewel for jewel in jewels}
        count = 0
        for curr in stones:
            if curr in jewels:
                count += 1
        return count