class Solution:
    def lengthOfLongestSubstringKDistinct(self, s: str, k: int) -> int:
        """
        The idea here is we have a string and we want to maintain a window of size k
        
        eceba
        s
        e
        while len(hashtable) is greater than k we want to move forward start until our window == k
        
        
        else:
        we want to move forward end and check it to a max value
        
         Given a string s and an integer k, return the length of the longest substring of s that contains at most k distinct characters.

 

Example 1:

Input: s = "eceba", k = 2
Output: 3
Explanation: The substring is "ece" with length 3.
Example 2:

Input: s = "aa", k = 1
Output: 2
Explanation: The substring is "aa" with length 2.
 
        
        
        """
        
        hashtable = {}
        start = 0
        maximum = 0
        for end in range(len(s)):
            while len(hashtable) > k:
                
                # remove the values from the hashtable
                if s[start] in hashtable:
                    hashtable[s[start]] -= 1
                    if hashtable[s[start]] == 0:
                        del hashtable[s[start]]
                    start += 1
            if s[end] in hashtable:
                hashtable[s[end]] +=1
            else:
                hashtable[s[end]] = 1
            if len(hashtable) <= k:
                maximum = max(maximum, end-start + 1)
                
        return maximum
        
        
        