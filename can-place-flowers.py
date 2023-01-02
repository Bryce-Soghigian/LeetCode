class Solution:
    def _check_left(self, flowerbed, i):
        if i-1 >= 0 and flowerbed[i-1] == 0:
            return True
        return False
    def _check_right(self, flowerbed, i):
        if i+1 <= len(flowerbed) -1 and flowerbed[i+1] == 0:
            return True
        return False
        
    def _valid_flower_state(self, flowerbed, i):
        if i == 0:
            return self._check_right(flowerbed, i)
        elif i == len(flowerbed) - 1:
            return self._check_left(flowerbed, i)
        else:
            return self._check_left(flowerbed, i) and self._check_right(flowerbed, i)
            
        
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        if len(flowerbed) == 1 and flowerbed[0] == 0:
            n -= 1
            return n <= 0
            
        
        for i, val in enumerate(flowerbed):
            if val == 0:
                # check if curr_state is a valid place to put a flower
                if self._valid_flower_state(flowerbed, i):
                    n -= 1
                    flowerbed[i] = 1
            if n == 0:
                return True
        return n <= 0
