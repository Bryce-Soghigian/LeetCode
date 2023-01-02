var findAndReplacePattern = function(words, pattern) {
  let ret = [];
  let matching;
  for (let elm of words) {
    matching = true;
    for (let i = 0, pat={}; i<elm.length && matching; i++) {
      if (pat[pattern[i]]) { // if our pat hash has already been set for chat pattern[i]
        if (pat[pattern[i]] != elm[i]) { // then it must be equal current letter
          matching = false;  // if not, we don't have a match
        }
      } else if (Object.values(pat).includes(elm[i])){ // tricky part, check that the value has not already been set
        matching = false;
      } else {
        pat[pattern[i]] = elm[i];
      }
    }
    if (matching) ret.push(elm);
  }
  return ret;
};