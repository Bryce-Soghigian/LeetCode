
var stringShift = function(s, shift) {
  // turn s into an array and perform shifts that way
  // calculate the number of shifts necessary - left and right cancel each other out
  // perform shifts necessary
  let totalShift = 0;
  shift.forEach(move => {
    // if left, subtract amount of moves
    // if right, add amount of moves
    // if end amounts is positive, shift right that amount
    // if end amounts is negative, shift left that amount
    if (move[0] === 0) {
      totalShift -= move[1];
    } else {
      totalShift += move[1];
    }
  });
  sArr = s.split("");
  if (totalShift > 0) {
    // move right - last char to beginning
    for (let i = 0; i < totalShift; i++) {
      let lastChar = sArr.pop();
      sArr.unshift(lastChar);
    }
  }
  if (totalShift < 0) {
    // shift left - first char to end
    for (let i = 0; i < totalShift * -1; i++) {
      let firstChar = sArr.shift();
      sArr.push(firstChar);
    }
  }
  return sArr.join("");
}