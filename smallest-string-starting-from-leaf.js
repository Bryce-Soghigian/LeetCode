const smallestFromLeaf = (root, parent = '') => {
    if (root === null) return parent;
    const char = String.fromCharCode(root.val + 97);
    if (root.right === null) return smallestFromLeaf(root.left, char + parent);
    if (root.left === null) return smallestFromLeaf(root.right, char + parent);
    
    const leftSmallest = smallestFromLeaf(root.left, char + parent);
    const rightSmallest = smallestFromLeaf(root.right, char + parent);
    return leftSmallest < rightSmallest ? leftSmallest : rightSmallest;      
};