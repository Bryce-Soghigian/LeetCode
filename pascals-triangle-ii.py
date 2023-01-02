class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        ans=[[1]]
        def prepare_row(index,ans):
            if index==rowIndex:
                return ans
            else:
                ans.append([])
                index+=1
                for i in range(1,index):
                    ans[index].append(ans[index-1][i-1]+ans[index-1][i])
                ans[index].append(1)
                ans[index].insert(0,1)
            prepare_row(index,ans)
        prepare_row(0,ans)
        return ans[-1]