class Solution:
    def candy(self, ratings: List[int]) -> int:
        left_to_right = [1 for i in range(len(ratings))]
        right_to_left = [1 for i in range(len(ratings))]
        
        for i in range(1,len(ratings)):
            if ratings[i] > ratings[i-1]:
                left_to_right[i] = left_to_right[i-1] + 1
        
        for i in range(len(ratings)-2, -1,-1,):
            if ratings[i] > ratings[i+1]:
                right_to_left[i] = right_to_left[i+1] + 1
        print(left_to_right,right_to_left)
        total = 0
        for i in range(len(ratings)):
            total += max(right_to_left[i],left_to_right[i])
        return total