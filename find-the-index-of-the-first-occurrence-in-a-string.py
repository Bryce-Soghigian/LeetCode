class Solution:
    def strStr(self, haystack, needle):
        hlen, nlen = len(haystack), len(needle)
        if nlen == 0:
            return 0
        if nlen > hlen or hlen == 0:
            return -1

        rolling = lambda x, y: x * 29 + y
        get_hash = lambda ch: ord(ch) - ord('a')
        
        nhash = reduce(rolling, map(get_hash, needle))
        hhash = reduce(rolling, map(get_hash, haystack[:nlen]))
        if nhash == hhash:
            return 0

        high_base = 29 ** (nlen - 1)
        for i in range(nlen, hlen):
            hhash -= get_hash(haystack[i - nlen]) * high_base # remove first in hash code
            hhash = rolling(hhash, get_hash(haystack[i]))     # add new
            if nhash == hhash:
                return i - nlen + 1

        return -1