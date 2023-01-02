var addToArrayForm = function(arr, k) {
    k = `${k}`.split('');
    var carry = 0;
    var arrIndex = arr.length - 1;
    var kIndex = k.length - 1;

    var ret = [];
    while (arrIndex >= 0 || kIndex >= 0 || carry > 0) {
        var cur = carry;
        cur += arrIndex >= 0 ? arr[arrIndex] : 0;
        cur += kIndex >= 0 ? Number.parseInt(k[kIndex]) : 0;

        ret.push(cur % 10);
        carry = Math.floor(cur / 10);

        arrIndex -= 1;
        kIndex -= 1;
    }

    return ret.reverse();
};