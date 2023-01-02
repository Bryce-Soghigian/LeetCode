class Solution(object):
    def countShips(self, sea: 'Sea', topRight: 'Point', bottomLeft: 'Point') -> int:
        
        def helper(x1,x2,y1,y2):
            if not (x1<=x2 and y1<=y2 and sea.hasShips(Point(x2,y2),Point(x1,y1))): return 0
            if x1==x2 and y1==y2: return 1
            midx=(x1+x2)//2
            midy=(y1+y2)//2
            total=0
            total+=helper(x1,midx,y1,midy)
            total+=helper(x1,midx,midy+1,y2)
            total+=helper(midx+1,x2,midy+1,y2)
            total+=helper(midx+1,x2,y1,midy)
            return total
            
        return helper(bottomLeft.x,topRight.x,bottomLeft.y,topRight.y)