class CustomStack:

    def __init__(self, maxSize: int):
        self.max_size = maxSize
        self.stack = []
        
        

    def push(self, x: int) -> None:
        if len(self.stack) != self.max_size:
            self.stack.append(x)
            

    def pop(self) -> int:
        if len(self.stack) == 0:
            return -1
        return self.stack.pop()
        

    def increment(self, k: int, val: int) -> None:
        
        if k > len(self.stack):
            
            for i in range(len(self.stack)):
                self.stack[i] += val
        else:
            
            for i in range(k):
                self.stack[i] += val
        


# Your CustomStack object will be instantiated and called as such:
# obj = CustomStack(maxSize)
# obj.push(x)
# param_2 = obj.pop()
# obj.increment(k,val)