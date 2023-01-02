from collections import deque
class TrieNode:
    def __init__(self,letter):
        self.letter = letter
        self.isWord = False
        self.children = {}

    
class WordDictionary:
    def __init__(self):
        self.root = TrieNode("")

    def addWord(self, word):
        curr = self.root

        for letter in word:
            if letter in curr.children:
                curr = curr.children[letter]
            else:
                new_trie_node = TrieNode(letter)
                curr.children[letter] = new_trie_node
                curr = new_trie_node
        
        curr.isWord = True
        

    def search(self, word):
        """Write a bfs theen sheesh"""
        queue = deque()
        queue.append((0,self.root))
        while len(queue) != 0:
            index,node = queue.popleft()
            if index > len(word):
                break
            if index == len(word):
                if node.isWord:
                    return True
                continue
            for next_letter,next_node in node.children.items():
                if word[index] in (".",next_letter):
                    queue.append((index+1,next_node))

        #Get through whole traversal and havent returned true its invalid sadge
        return False

        


# Your WordDictionary object will be instantiated and called as such:
# obj = WordDictionary()
# obj.addWord(word)
# param_2 = obj.search(word)