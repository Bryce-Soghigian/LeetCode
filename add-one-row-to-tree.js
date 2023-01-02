var addOneRow = function(root, v, d) {
    
    function callDFS(node, depth, dir) {
        if(depth === d) {
            if(dir === 'L') return new TreeNode(v, node, null)
            else return new TreeNode(v, null, node)
        }
        if(!node) return null;
        
        node.left = callDFS(node.left, depth+1, 'L');
        node.right = callDFS(node.right, depth+1, 'R');
        return node;
    } 
    return callDFS(root, 1, 'L')
};