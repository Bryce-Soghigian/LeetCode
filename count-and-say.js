const getSay = s => {
    let output = ''
    let count = 1
    let prev = s[0]
    if (s.length === 1) return `${count}${prev}` 
    for (let i = 1; i < s.length; i++) {
        let ch = s[i]
        if (ch !== prev) {
            output += `${count}${prev}`
            prev = ch
            count = 1
        } else {
            count++            
        }
        if (i === s.length -1 ) {
            output += `${count}${prev}`
        }
    }
    return output
}

const memo = {}
const countAndSay = n => { 
    let next = '1'
    if (memo[n]) return memo[n]
    for (let i = 1; i < n; i++) {
        if (memo[i+1]) {
            next = memo[i+1]
            continue
        }
        memo[i] = next
        next = getSay(next)
    }
    memo[n] = next
    return next
};
