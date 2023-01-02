class Solution:
    def fourSum(self, nums, target):
        nums.sort()
        found = set()
        st = defaultdict(list)
        cnt = i = 0
        lastVal = None
        while (i < len(nums)-2):
            if lastVal == nums[i]:  # track duplicates
                    cnt +=1
                    if cnt >2:  # outer loop has seen value two times (if there are more repeats inner loop would have seen them)
                        i +=1
                        continue
            else: 
                lastVal, cnt = nums[i], 1
            for j in range(i+1,len(nums)):
                sm = nums[i] + nums[j]
                comp = target - sm
                if comp in st:
                    for k, l in st[comp]:
                        if i!=k and i!=l and j!=k and j!=l:
                            hit = [nums[i],nums[j],nums[k],nums[l]]
                            hit.sort()
                            found.add(tuple(hit))
                st[sm].append((i,j))
            i+=1
        return list(found)