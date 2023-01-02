/**
 *
 * @param {string} w1
 * @param {string} w2
 */
const getOverlap = (w1, w2) => {
  for (var i = 0; i < w1.length; i++) {
    if (w1.substring(i) === w2.substring(0, w1.length - i)) {
      return w1.length - i;
    }
  }
  return 0;
};

/**
 * @param {string[]} words
 * @return {string}
 */
var shortestSuperstring = function (words) {
  const matrix = [];

  for (var i = 0; i < words.length; i++) {
    matrix.push([]);
    for (var j = 0; j < words.length; j++) {
      matrix[i][j] = getOverlap(words[i], words[j]);
    }
  }

  //   const memo = {};
  const memo = []
  for (var i = 0; i < words.length; i++) {
    memo.push([]);
  }
  getMaxOverlap = (last, mask) => {
    if (memo[last][mask]) return memo[last][mask];
    if (mask === 0) return 0;
    let max = -1;
    for (var i = 0; i < words.length; i++) {
      if ((mask & (1 << i)) == 0) continue;
      const attempt =
        (matrix?.[last]?.[i] || 0) + getMaxOverlap(i, mask - (1 << i));

      if (attempt > max) {
        max = attempt;
      }
    }
    memo[last][mask] = max;
    return memo[last][mask];
  };
  let max = 0;
  let bestWord = 0;
  const fullMask = (1 << words.length) - 1;
  for (var i = 0; i < words.length; i++) {
    const maxOverlap = getMaxOverlap(i, fullMask - (1 << i));
    if (maxOverlap > max) {
      max = maxOverlap;
      bestWord = i;
    }
  }

  let ret = words[bestWord];
  let mask = fullMask - (1 << bestWord);

  while (mask) {
    let max = -1;
    let newBest = "";
    for (var i = 0; i < words.length; i++) {
      if ((mask & (1 << i)) == 0) continue;
      const val =
        (matrix?.[bestWord]?.[i] || 0) + getMaxOverlap(i, mask - (1 << i));
      if (val > max) {
        max = val;
        newBest = i;
      }
    }
    ret += words[newBest].substring(matrix[bestWord][newBest]);
    bestWord = newBest;
    mask -= 1 << bestWord;
  }
  return ret;
};