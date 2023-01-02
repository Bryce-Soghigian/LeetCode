class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        word_set = set(wordDict)
        dp = [False] * (len(s) + 1)
        dp[0] = True

        for i in range(1, len(s) + 1):
            for j in range(i):
                if dp[j] and s[j:i] in word_set:
                    dp[i] = True
                    break
        return dp[len(s)]
# class TrieNode:
#     def __init__(self, val):
#         self.val = val
#         self.is_word = False
#         self.children = {}
# class Solution:
    
    
#     def _add_word(self, word):
#         curr = self.trie
#         for char in word:
#             if char not in curr.children:
#                 new_trie_node = TrieNode(char)
#                 curr.children[char] = new_trie_node
#             curr = curr.children[char]
#         curr.is_word = True
        
#     def _build_trie(self, word_list):
#         for word in word_list:
#             self._add_word(word)
            
#     def wordBreak(self, s: str, wordDict: List[str]) -> bool:
#         """
#         aplepenapple
#             s
#         s
#             p   a 
#            e     p 
#          n        p 
#                    l 
#                     e 
        
#         1. Build a trie
#         2. iterate through our list.
#         2a check to see if our current iterator is a child of our current trie node
#         2b if true:
#             advance pointer
#         2c if false:
#             return False cant make word
#             if trie pointer is at is_word, set trie pointer back to root
        
        
# "aaaaa
# aa"
# ["aaaa","aa"]
#         """
#         self.trie = TrieNode("*")
#         trie_pointer = self.trie
#         self._build_trie(wordDict)
#         for i, char in enumerate(s):
#             if char not in trie_pointer.children:
#                 return False
#             else:
#                 trie_pointer = trie_pointer.children[char]
#                 if trie_pointer.is_word and i != len(s) - 1:
#                     trie_pointer = self.trie
#         if trie_pointer.is_word:
#             return True
                    
#         return False
        