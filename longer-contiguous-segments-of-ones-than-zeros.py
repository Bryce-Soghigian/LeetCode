class Solution:
    def checkZeroOnes(self, s: str) -> bool:
        longest_chain_of_ones = 0
        longest_chain_of_zeros = 0
        zeroes = 0
        ones = 0
        
        
        for i in range(len(s)):
            if s[i] == "1":
                ones += 1
                zeroes = 0
            else:
                zeroes += 1
                ones = 0
            longest_chain_of_ones = max(longest_chain_of_ones, ones)
            longest_chain_of_zeros = max(longest_chain_of_zeros, zeroes)
        return longest_chain_of_ones > longest_chain_of_zeros
