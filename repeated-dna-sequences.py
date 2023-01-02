class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        """
        Maintain two sets.
        
        As we see one dna sequence if its been seen before we want to append it to output set
        what is a dna-sequenxce? its a substring of s[i: i + 10]
        """
        output = set()
        unique = set()
        
        for i in range(len(s)):

            new_str = s[i: i + 10]
            if new_str in unique:
                output.add(new_str)
            unique.add(new_str)
        return list(output)