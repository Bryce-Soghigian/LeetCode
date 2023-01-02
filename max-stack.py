class MaxStack:

    def __init__(self):
        self.stack = []
        
    def push(self, x: int) -> None:
        self.stack.append(x)
        
    def pop(self) -> int:
        return self.stack.pop()

    def top(self) -> int:
        return self.stack[-1]

    def peekMax(self) -> int:
        return max(self.stack)

    def popMax(self) -> int:
        maxx = max(self.stack)
        for i in range(len(self.stack)-1, -1, -1):
            if self.stack[i] == maxx:
                return self.stack.pop(i)