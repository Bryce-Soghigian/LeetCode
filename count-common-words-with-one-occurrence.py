class Solution:
    def countWords(self, words1: List[str], words2: List[str]) -> int:
        dct1, dct2  = collections.Counter(words1), collections.Counter(words2)
        new1 = {el:freq for el, freq in dct1.items() if freq == 1}
        new2 = {el:freq for el, freq in dct2.items() if freq == 1}      
        intrsct = new1.keys() & new2.keys()
        return len(intrsct)