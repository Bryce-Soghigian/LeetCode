class Solution:
    def guessNumber(self, n: int) -> int:
        low=1
        high=n
        while low<=high:
            mid=(low+high)>>1
            val=guess(mid)
            if val==-1:
                high=mid-1
            elif val==1:
                low=mid+1
            else:
                return mid