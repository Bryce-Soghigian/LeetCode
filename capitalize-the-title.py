class Solution:
    def capitalizeTitle(self, title: str) -> str:
        title = title.lower()
        
        paragraph = title.split(" ")
        output = ""
        for word in paragraph:
            if len(word) <= 2:
                output += word 
            else:
                output += word[0].upper() + word[1:]
            output += " "
        return output[:-1]