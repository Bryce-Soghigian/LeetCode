class Solution:
    def diStringMatch(self, S: str) -> List[int]:
        ARRR=[i for i in range(len(S)+1)]
        return [ARRR.pop((j=='I')-1) for j in S]+ARRR
        