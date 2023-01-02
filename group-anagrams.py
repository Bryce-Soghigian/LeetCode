class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagrams = defaultdict(list)
        for item in strs:
            sorted_item = "".join(sorted(item))
            anagrams[sorted_item].append(item)
        return_arr = list(anagrams.values())
        return return_arr
	