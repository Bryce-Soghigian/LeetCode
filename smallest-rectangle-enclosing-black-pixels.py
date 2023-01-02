class Solution:
    def minArea(self, image: List[List[str]], x: int, y: int) -> int:
        yy = len(image[0])
        x_min = yy
        x_max = 0
        flag = 0
        delete = -1
        
        for index1,list1 in enumerate(image):
            counting = list1.count('1')
            if counting != 0:
                flag = 1
                x_min = min(x_min,list1.index('1'))
                x_max = max(x_max, yy - list1[::-1].index('1') - 1)                
            else:
                if flag == 1:
                    return((x_max - x_min+1)*(index1 - delete-1))
                else:
                    delete += 1
        
        return((x_max - x_min+1)*(index1 - delete))