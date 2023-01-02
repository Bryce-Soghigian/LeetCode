class Solution:
    def longestDiverseString(self, a: int, b: int, c: int) -> str:
        
        # Get letters in reverse order by their count
        letters, cnt = 'abc', [a, b, c]
        d = list(sorted(zip(letters, cnt), key=lambda x: x[1], reverse=True))

        # Place the most common letter as a row sequence, 
        # min() guarantees that other letters can make the string happy
        res = d[0][0] * (min(d[0][1], 2 * (1 + d[1][1] + d[2][1])))
        
        # Form array of other letters - it's guarantted we can place them, 
        # as the longest one is already in place
        lts = ''
        for idx, el in enumerate(d):
            if idx > 0: lts += el[0] * el[1]
         
        # Place other letters: firstly place a letter between each pair of the most common,
        # if reached end of string - return and form pairs
        i, shift = 2, 2
        for c in lts:
            if i > len(res)+1: 
                shift += 1
                i = shift
            res = res[:i] + c + res[i:]
            i += shift + 1

        return res