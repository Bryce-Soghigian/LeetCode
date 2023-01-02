class Solution:
    def starts_with_vowel(self,word):
        word = word.lower()
        vowels = {'a','e','i','o','u'}
        if word[0] in vowels:
            return True
        else:
            return False
    def toGoatLatin(self, sentence: str) -> str:
        output = ""
        sentence = sentence.split(sep=" ")
        
        for i in range(len(sentence)):
            if self.starts_with_vowel(sentence[i]):
                sentence[i] += "ma"
            else:
                # Remove first char append it to the end and add ma
                first = sentence[i][0]
                sentence[i] = sentence[i][1:] + first + "ma"
            sentence[i] += ("a" * (i + 1))
            output += sentence[i]
            if i != len(sentence) -1:
                output += " "
        
        return output
                
                
        
        