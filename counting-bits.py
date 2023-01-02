class Solution:
    def get_bits(self, binary_number):
        bits = 0
        for i in range(2,len(binary_number)):
            bit = binary_number[i]
            if bit == "1":
                bits += 1
        return bits
    def countBits(self, n: int) -> List[int]:
        
        return [self.get_bits(bin(curr)) for curr in range(n+1)]
        