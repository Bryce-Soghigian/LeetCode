class Solution:
    def expand(self, s: str) -> List[str]:

        res = [""]
        i = 0

        while( i<len(s) ):
            if( s[i]=="{" ):
                temp = []
                prefix = res
                while( s[i]!="}" ):
                    if( s[i].isalpha() ):
                        temp += [ss+s[i] for ss in prefix]
                    i += 1  
                res = sorted(temp)

            elif( s[i].isalpha() ):
                res = [ss+s[i] for ss in res]

            i += 1

        return res