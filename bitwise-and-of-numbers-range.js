var rangeBitwiseAnd = function(m, n) {
    // in order to not be 0, n and m must have the same amount of bits
    // e.g. 2 and 6 will be 0 since its 10 and 110.
    // e.g. 5 and 7 will not be 0 since its 101 and 111
    let mBin = m.toString(2).split('')
    let nBin = n.toString(2).split('')
    if ( mBin.length !== nBin.length ) {
        return 0
    }
    else {
        // to get from m to n, you must turn on/off each bit that isn't used in both m and n
        // turning on and off with negate those bits out, so only need to find bits shared by m and n
        // right shift until equal, count shifts, shift back
        let shifts = 0
        while ( m !== n ) {
            m = m >> 1
            n = n >> 1
            shifts += 1
        }
        return m << shifts
    }
};