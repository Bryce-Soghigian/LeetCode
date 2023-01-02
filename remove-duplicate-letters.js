/**
 * @param {string} s
 * @return {string}
 */
const map = {
    a: 0,
    b: 1,
    c: 2,
    d: 3,
    e: 4,
    f: 5,
    g: 6,
    h: 7,
    i: 8,
    j: 9,
    k: 10,
    l: 11,
    m: 12,
    n: 13,
    o: 14,
    p: 15,
    q: 16,
    r: 17,
    s: 18,
    t: 19,
    u: 20,
    v: 21,
    w: 22,
    x: 23,
    y: 24,
    z: 25,
}

var removeDuplicateLetters = function(s) {
    if (!s) {
        return ''
    }
    
    const keys = Array.from(new Set(s.split('')))
    
    keys.sort((a,b) => map[a] - map[b])
    
    return walk(s, keys)
};

function walk (s, keys) {
    if (!s || !keys.length) {
        return ''
    }
    
    for (let i = 0, len = keys.length; i < len; i++) {
        const sub = s.substring(s.indexOf(keys[i]))
        const tmp = Array.from(new Set(sub.split('')))
        
        if (tmp.length !== keys.length) {
            continue
        }
        const copy = keys.slice()
        const pop = copy.splice(i, 1)
        
        const str = sub.substring(1).replace(new RegExp(pop, 'g'), '')
        
        const res = walk(str, copy)
        return sub[0] + res
    }
}