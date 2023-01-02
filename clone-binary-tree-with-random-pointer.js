/**
 * // Definition for a Node.
 * function Node(val, left, right, random) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.random = random === undefined ? null : random;
 * };
 */

/**
 * @param {Node} root
 * @return {NodeCopy}
 */
var copyRandomBinaryTree = function(root) {
        if (!root) return null
    const clone_root = new NodeCopy()
    let arr = [root]
    const map = new Map()
    while (arr.length>0) {
        const node = arr.shift()
        map.set(node, new NodeCopy(node.val))
        node.left && arr.push(node.left)
        node.right && arr.push(node.right)
    }
    function copy(node) {
        if (node == null) return null
        const clone_node = map.get(node)
        clone_node.random = map.get(node.random)
        clone_node.left = copy(node.left)
        clone_node.right = copy(node.right)
        return clone_node
    }
    return copy(root)
};