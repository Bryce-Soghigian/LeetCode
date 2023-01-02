class Solution:
    def isUgly(self, n: int) -> bool:
        if n <=0:
            return False
        if n == 1 or n==3 or n == 5:
            return True
        else:
            while True:
                if n %5==0:
                    n/=5
                elif n%3 == 0:
                    n/=3
                elif n%2 == 0:
                    n/=2
                else:
                    if n == 1:
                        return True
                    else:
                        return False