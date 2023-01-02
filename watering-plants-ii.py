class Solution(object):
    def minimumRefill(self, plants, capacityA, capacityB):
        """
        :type plants: List[int]
        :type capacityA: int
        :type capacityB: int
        :rtype: int
        """
        bucket1 = capacityA
        bucket2 = capacityB
        idx = 0
        jdx = len(plants)-1
        total = 0
        while idx < jdx:
            bucket1 = bucket1-plants[idx]
            if bucket1<0:
                bucket1=capacityA-plants[idx]
                total+=1
            bucket2 = bucket2 - plants[jdx]
            if bucket2<0:
                bucket2=capacityB - plants[jdx]
                total+=1
            idx+=1
            jdx-=1
        if jdx == idx:
            if bucket1<bucket2:
                bucket2 = bucket2-plants[idx]
                if bucket2<0:
                    bucket2=capacityB-plants[idx]
                    total+=1
            else:
                bucket1 = bucket1-plants[idx]
                if bucket1<0:
                    bucket1=capacityA-plants[idx]
                    total+=1
        return total