const utils = require('../../utils.ts');

// Grab the input
// let input = utils.inputString('input.txt');
let input = utils.inputString('test.txt');
let inputStrArray = utils.inputStringArray(input);
let inputArray = utils.stringArrayToNumber(inputStrArray);

function step1() {
  let joltMap = new Map();
  let maxJolt = 0;
  inputArray.forEach((n: number) => {
    if (n >= maxJolt) {
      maxJolt = n;
    }
    if (joltMap.has(n)) {
      console.log('Duplicate');
      return;
    }
    joltMap.set(n, 0);
  });

  let numOne = 0;
  let numThree = 1;
  let i = 0;
  for (i = 0; i < maxJolt; ) {
    // Check for 1 jolt diff
    if (joltMap.has(i + 1)) {
      i++;
      numOne++;
      continue;
    }

    // Check for 2 jolt diff
    if (joltMap.has(i + 2)) {
      i += 2;
      continue;
    }
    // Check for 3 jolt diff
    if (joltMap.has(i + 3)) {
      i += 3;
      numThree++;
      continue;
    }
  }

  return numOne * numThree;
}

function step2() {
  let joltMap = new Map();
  let maxJolt = 0;
  inputArray.forEach((n: number) => {
    if (n >= maxJolt) {
      maxJolt = n;
    }
    if (joltMap.has(n)) {
      console.log('Duplicate');
      return;
    }
    joltMap.set(n, 0);
  });

  let combos = 1;
  let i = 0;
  for (i = 0; i < maxJolt; ) {
    let addi = 0;
    let addCombo = 0;
    // Check for 3 jolt diff
    if (joltMap.has(i + 3)) {
      addCombo += 1;
      addi = 3;
    }

    // Check for 2 jolt diff
    if (joltMap.has(i + 2)) {
      addCombo += 1;
      addi = 2;
    }
    // Check for 1 jolt diff
    if (joltMap.has(i + 1)) {
      addCombo += 1;
      addi = 1;
    }
    console.log(addCombo);
    i += addi;
    combos += addCombo;
  }
  return combos;
}

console.log(step1());
console.log(step2());

console.log('end');
