# class TrieNode:
#     def __init__(self, value, is_word=False):
#         self.value = value
#         self.is_word = is_word
#         self.children = {}
    
# class Trie:
#     def __init__(self):
#         self.root = TrieNode("*")
        
#     def add(self, word):
        
#         curr = self.root
#         for char in word:
#             if char not in curr.children:
#                 # define a new child node
#                 child = TrieNode(char)
#                 curr.children[char] = child
#             curr = curr.children.get(char)
        
#         curr.is_word = True
        
# class StreamChecker:
#     def __init__(self, words: List[str]):
#         self.root = Trie()
#         self.stack = deque()
#         for word in words:
#             self.root.add(reversed(word))

#     def query(self, letter: str) -> bool:
#         self.stack.append(letter)
#         pointer = len(self.stack) - 1
#         curr = self.root.root
#         while pointer >= 0:
#             if curr.is_word:
                
#                 return True
#             if self.stack[pointer] in curr.children:
#                 curr = curr.children[self.stack[pointer]]
#             else:
#                 return False
            
#             pointer -= 1
class StreamChecker:

    def __init__(self, words: List[str]):
        self.trie = {}
        self.stream = deque([])

        for word in set(words):
            node = self.trie       
            for ch in word[::-1]:
                if not ch in node:
                    node[ch] = {}
                node = node[ch]
            node['$'] = word
        
        
    def query(self, letter: str) -> bool:
        self.stream.appendleft(letter)
        
        node = self.trie
        for ch in self.stream:
            if '$' in node:
                return True
            if not ch in node:
                return False
            node = node[ch]
        return '$' in node