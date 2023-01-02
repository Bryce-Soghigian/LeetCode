class OrderedStream:

    def __init__(self, n: int):
        self.order = [None] * n
        self.pt = 0  # 0-indexed 

    def insert(self, idKey: int, value: str) -> List[str]:
        idKey -= 1  # 0-indexed 
        self.order[idKey] = value
        if idKey > self.pt: return []  # not reaching pt 

        while self.pt < len(self.order) and self.order[self.pt]: self.pt += 1  # update self.pt 
        return self.order[idKey:self.pt]