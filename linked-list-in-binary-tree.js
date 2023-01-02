const isSubPath = (head, root) => {
    const queue = [root];
    
    while(queue.length > 0) {
        const curr = queue.pop();
        
        if(helper(curr, head)) return true;
        
        if(curr.left) queue.unshift(curr.left);
        if(curr.right) queue.unshift(curr.right);
    }
    
    return false;
    
};

const helper = (treeNode, listNode) => {
    if(!treeNode) return false;
    
    if(treeNode.val !== listNode.val) return false;
    
    if(!listNode.next) return true;
    
    listNode = listNode.next;
    
    const leftResult = helper(treeNode.left, listNode);
    const rightResult = helper(treeNode.right, listNode);
    
    return leftResult || rightResult;
};