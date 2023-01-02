const assert = require('assert');

/**
 * @param {number[][]} courses
 * @return {number}
 */
var scheduleCourse = function(courses) {
  const sorted = courses.sort((a,b) => a[1] - b[1]);

  const learned = [];
  let day = 0;

  sorted.forEach(([duration, endDay]) => {
    add(duration);
    day+=duration;
    if (day > endDay) {
      day-=remove();
    }
  });

  return learned.length;

  function add(val) {
    let index = learned.length;
    let parent = getParent(index);
    learned.push(val);
    while(index > 0 && learned[index] > learned[parent]) {
      swap(index, parent);
      index = parent;
      parent = getParent(index);
    }
  }

  function remove() {
    swap(0, learned.length - 1);
    const res = learned.pop();
    let index = 0;
    let [left, right] = getChildren(index);
    while(index < learned.length && (learned[index] < learned[left] || learned[index] < learned[right])) {
      const largerChild = learned[left] < learned[right] ? right : left;

      swap(index, largerChild);

      index = largerChild;
      [left, right] = getChildren(index);
    }
    return res;
  }

  function swap(a, b) {
    const d = learned[a];
    learned[a] = learned[b];
    learned[b] = d;
  }

  function getParent(index) {
    return Math.floor((index - 1) / 2);
  }

  function getChildren(index) {
    return [2 * index + 1, 2*index + 2];
  }
};


