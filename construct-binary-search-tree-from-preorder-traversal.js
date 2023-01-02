var bstFromPreorder = function(preorder) {
    let root = new TreeNode(preorder[0])
    for (let i = 1; i < preorder.length; i++) {
        appendToTreeNode(root, preorder[i])
    }
    return root
};

function appendToTreeNode(root, val) {
    if (val <= root.val) {
        if (root.left) appendToTreeNode(root.left,val);
        else root.left = new TreeNode(val);
    } else {
        if (root.right) appendToTreeNode(root.right,val);
        else root.right = new TreeNode(val);
    }
}