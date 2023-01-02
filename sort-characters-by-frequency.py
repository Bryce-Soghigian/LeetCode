class Solution:
    def frequencySort(self, s: str) -> str:
        return "".join([tup[0] * tup[1] for tup in sorted(Counter(s).items(), key=lambda a: a[1], reverse=True)])
