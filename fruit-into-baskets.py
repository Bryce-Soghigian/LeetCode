class Solution:
    def totalFruit(self, tree: List[int]) -> int:
        ## Sliding window where we can collect at most two elements 
        k = 2
        hash_map = {}
        maximum_window = 0
        start = 0
        end = 0
        while end != len(tree):
            # sheeesh we gotta move forward our window
            if len(hash_map) == k and tree[end] not in hash_map:
                #move forward start until we can move forward end
                while len(hash_map) == k:
                    if tree[start] in hash_map:
                        if hash_map[tree[start]] == 1:
                            del hash_map[tree[start]]
                        else:
                            hash_map[tree[start]] -= 1
                        start+= 1
            else:
                # add
                if tree[end] in hash_map:
                    hash_map[tree[end]] += 1
                else:
                    hash_map[tree[end]] = 1
                maximum_window = max(maximum_window,end-start + 1)
                end+= 1
        return maximum_window
            
                        
                
        