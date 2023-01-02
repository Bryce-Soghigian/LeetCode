class Solution:
    def removeStones(self, stones):
        UF = {}
        def find(x):
            if x != UF[x]:
                UF[x] = find(UF[x])
            return UF[x]
        
        def union(x, y):
            print(x,"x")
            print(y, "y")
            if x not in UF:
                UF[x] = x
            if y not in UF:
                UF[y] = y
            rootX = find(x)
            rootY = find(y)
            UF[rootX] = rootY
        maxX = 10**4
        for x,y in stones:
            union(x,y+maxX)
        
        ### Finally, we go through each element in UF and find its root, count how many connected components (unique roots) are there and subtract it from the total number of stones.
        return len(stones) - len({find(n) for n in UF})