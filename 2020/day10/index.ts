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
  let numThree = 0;
  let i = 0;
  for (i = 0; i <= maxJolt; ) {
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
    if (i == maxJolt) {
      numThree++;
      break;
    }
  }
  return numOne * numThree;
}

// Memoize step 2. Remember the split points and go back to them as a way of
// confirming the new branch.
//
// i.e. if there is a +1, +2, +3 option at jolt 4, continue on with +1 then go
// back to 4 and continue with +2, then go back and continue with +3. If the
// memoization is in reverse order, we can then immediately sum up the potential
// branches from there.
//
// So first iteration through creates a map that is map[volt]potentialBranches
// [4,3] [5,2] etc
//
// should be then able to convert potential to valid in reverse order?
//
// Thing memoized fibinooci, draw out the tree that are the potential branches
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
  joltMap.set(maxJolt + 3, 0);
  console.log('max', maxJolt);
  let combos = 1;
  let i = 0;
  // Try going in reverse order by starting at maxJolt and memoizing the ways to
  // complete the circuit from there
  for (i = 0; i < maxJolt + 3; ) {
    let plusOne,
      plusTwo,
      plusThree = false;

    let addi = 0;
    // Check for 1 jolt diff
    if (joltMap.has(i + 1)) {
      plusOne = true;
      addi = 1;
    }
    // Check for 2 jolt diff
    if (joltMap.has(i + 2)) {
      plusTwo = true;
      addi = 2;
    }
    // Check for 3 jolt diff
    if (joltMap.has(i + 3)) {
      plusThree = true;
      addi = 3;
    }

    if (plusThree && plusTwo && plusOne) {
      combos *= 4;
    } else if (plusThree && plusTwo) {
      combos *= 2;
    } else if (plusThree && plusOne) {
      combos *= 2;
    } else if (plusTwo && plusOne) {
      combos *= 2;
    } else if (!plusThree && !plusTwo && !plusOne) {
      console.log('break', i);
      break;
    }
    i += addi;
  }
  return combos;
}

console.log(step1());
console.log(step2());

console.log('end');
