class TrieNode:
    def __init__(self, letter):
        self.letter = letter
        self.children = {}
        self.is_word = False

        
class Trie:
    def __init__(self):
        
        self.root = TrieNode(letter="*")
        self.longest_word = 0
        self.lcp = ""
    
    
    def _add_word(self, word):
        
        curr = self.root
        if len(word) == 0:
            curr.children[""] = TrieNode(letter="")
            curr.is_word = True
            return
        for letter in word:
            if letter not in curr.children:
                curr.children[letter] = TrieNode(letter=letter)
            curr = curr.children[letter]
        curr.is_word = True
        
    def add_strs(self, words):
        for word in words:
            self.longest_word = max(self.longest_word, len(word))
            self._add_word(word)
    
    def retrieve_longest_common_prefix(self):
        prefix = ""
        curr = self.root
        for _ in range(self.longest_word + 1):
            if curr.letter != "*":
                prefix += curr.letter
            if len(curr.children) > 1 or len(curr.children) == 0 or curr.is_word:
                self.lcp = prefix
                break
  
            curr = curr.children[[key for key in curr.children.keys()][0]]
            
        
        
        
        """
        
        letter = "*"
        children = {
        "f": TrieNode
        }
        
        queue.append(f)
        """
        

class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        """
        1. Create Trie
        2. BFS Until we find a level in trie that has more than one character
        3. Return prefix
        """
        trie = Trie()
        trie.add_strs(strs)
        trie.retrieve_longest_common_prefix()
        return trie.lcp