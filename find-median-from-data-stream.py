import bisect
class MedianFinder:
    """
    Find the median of a data stream in O(1).
    """
    def __init__(self):
        self.nums = []

    def addNum(self, num: int) -> None:
        """
        stream input method.
        """
        bisect.insort(self.nums, num)

    def findMedian(self) -> float:
        """
        Get median from self.nums.
        """
        if len(self.nums) % 2 == 0:
            # if even get the two inner most numbers and add them returning the average
            return (self.nums[len(self.nums) // 2 - 1] + self.nums[len(self.nums) // 2]) / 2
        else:
            # return the middlemost element in list
            return self.nums[len(self.nums) //2]