class MRUQueue:

    def __init__(self, n: int):
        self.nums = [x for x in range(1,n+1)]

    def fetch(self, k: int) -> int:
        self.nums.append(self.nums.pop(k-1))
        return self.nums[-1]