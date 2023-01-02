class Solution:

    # Generate all permutations of a string and try and combine them
    def arrayStringsAreEqual(self, word1: List[str], word2: List[str]) -> bool:
        word1 = "".join(word1)
        word2 = "".join(word2)
        return word1 == word2

        