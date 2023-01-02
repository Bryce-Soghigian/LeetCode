class Solution(object):
    def groupStrings(self, strings):
        """
        :type strings: List[str]
        :rtype: List[List[str]]
        """
        
        cache = collections.defaultdict(list)
        
        for s in strings:
            if len(s) == 1:
                cache[0].append(s)
            else:
                n = len(s)
                level = ""
                for i in range(1, n):
                    
                    diff = ord(s[i]) - ord(s[i-1])
                    if diff < 0:
                        diff += 26
                    level += str(diff)
                level += str(n)
                cache[level].append(s)
        
        return cache.values()