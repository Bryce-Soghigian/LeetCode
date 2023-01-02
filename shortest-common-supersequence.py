class Solution:
    def shortestCommonSupersequence(self, s: str, t: str) -> str:
        
        @cache
        def find_lcs(i,j):
            if i<0 or j<0:
                return ''
            
            if s[i]==t[j]:
                s1=find_lcs(i-1,j-1)
                return s1+s[i]
            
            else:
                s1=find_lcs(i,j-1)
                s2=find_lcs(i-1,j)

                return s1 if len(s1) >len(s2) else s2
            
        lcs=find_lcs(len(s)-1,len(t)-1)
        
        i,j,ans=0,0,''
        
        for c in lcs:
            
            while s[i] != c:
                ans += s[i]
                i += 1
                
            while t[j] != c:
                ans+= t[j]
                j += 1
                
            ans+=c
            i+=1
            j+=1
            
        return ans + s[i:] + t[j:]
        