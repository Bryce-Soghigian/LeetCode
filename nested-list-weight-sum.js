const depthSum = function (nestedList, depth = 1) {
    if (nestedList.length === 1 && nestedList[0].isInteger()) return nestedList[0].getInteger() * depth;

    let sum = nestedList.reduce((accum, cur) => {
        cur = (cur.isInteger()) ? cur.getInteger() * depth : depthSum(cur.getList(), depth + 1)
        return accum + cur;
    }, 0)

    return sum;
};