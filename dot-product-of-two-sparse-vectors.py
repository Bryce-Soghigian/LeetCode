class SparseVector:
    def __init__(self, nums: List[int]):
        self.vector = nums

    # Return the dotProduct of two sparse vectors
    def dotProduct(self, vec: 'SparseVector') -> int:
        product_vectors = vec.vector
        output = 0
        for i in range(len(product_vectors)):
            result = product_vectors[i] * self.vector[i]
            output += result
        return output

# Your SparseVector object will be instantiated and called as such:
# v1 = SparseVector(nums1)
# v2 = SparseVector(nums2)
# ans = v1.dotProduct(v2)