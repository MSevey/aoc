const utils = require('../../utils.ts');

// Grab the input
let input = utils.inputString('input.txt');
// let input = utils.inputString('test.txt');
let inputStrArray = utils.inputStringArray(input);
let inputArray = utils.stringArrayToNumber(inputStrArray);

let preAmble = 25;
let list = [];

function step1() {
  for (let n of inputArray) {
    // Build the initial list
    if (list.length < preAmble) {
      list.push(n);
      continue;
    }

    // Check if valid
    if (!isValid(n)) {
      return n;
    }

    // Update the list
    list.shift();
    list.push(n);
  }
}

function isValid(n: number) {
  let min = 0;
  let max = 0;
  list.forEach((a) => {
    list.forEach((b) => {
      if (a == b) {
        return;
      }
      let sum = a + b;
      if (min == 0 || sum < min) {
        min = sum;
      }
      if (max == 0 || sum > max) {
        max = sum;
      }
    });
  });
  if (n <= max && n >= min) {
    return true;
  }
  return false;
}

let sumList = [];
function step2(sum: number) {
  for (let n of inputArray) {
    sumList.push(n);
    let ls = listSum();
    if (ls < sum) {
      continue;
    }
    while (ls > sum && sumList.length > 2) {
      sumList.shift();
      ls = listSum();
    }
    if (ls == sum) {
      // return sum of first and last element
      return minMaxSum();
    }
  }
}

function minMaxSum() {
  let min = sumList[0];
  let max = sumList[0];
  sumList.forEach((n) => {
    if (n < min) {
      min = n;
    }
    if (n > max) {
      max = n;
    }
  });
  return min + max;
}

function listSum() {
  let sum = 0;
  sumList.forEach((n) => {
    sum += n;
  });
  return sum;
}

let s1res = step1();
console.log(s1res);
console.log(step2(s1res));

console.log('end');
