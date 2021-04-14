const utils = require('../../utils.ts');

// Grab the input
let input = utils.inputString('input.txt');
// let input = utils.inputString('test.txt');
let inputArray = utils.inputStringArray(input);

let accumulator = 0;

function step1(inputArray: string[]) {
  let instructionMap = new Map();
  let i = 0;
  for (i = 0; i < inputArray.length; ) {
    let uid = inputArray[i] + i.toString();
    if (instructionMap.has(uid)) {
      return;
    }
    instructionMap.set(uid, '');
    let elements = inputArray[i].split(' ');
    let instruction = elements[0];
    let value = parseInt(elements[1]);
    switch (instruction) {
      case 'acc':
        accumulator += value;
        i++;
        break;
      case 'jmp':
        i += value;
        break;
      default:
        i++;
    }
  }
}

step1(inputArray);
console.log('step1:', accumulator);

function indexes(inputArray: string[], ints: string) {
  let indexes = [];
  let i = 0;
  for (i = 0; i < inputArray.length; i++) {
    let instruction = inputArray[i].split(' ')[0];
    if (instruction == ints) {
      indexes.push(i);
    }
  }
  return indexes;
}

let nops = indexes(inputArray, 'nop');
let jmps = indexes(inputArray, 'jmp');

function isFix(inputArray: string[], indexes: number[], newInstr: string) {
  let fixed = true;
  let index = 0;
  for (index = 0; index < indexes.length; index++) {
    let j = indexes[index];
    accumulator = 0;
    let original = inputArray[j];
    let value = original.split(' ')[1];
    inputArray[j] = newInstr + ' ' + value;
    let instructionMap = new Map();
    let i = 0;
    fixed = true;
    for (i = 0; i < inputArray.length; ) {
      let uid = inputArray[i] + i.toString();
      if (instructionMap.has(uid)) {
        fixed = false;
        break;
      }
      instructionMap.set(uid, '');
      let elements = inputArray[i].split(' ');
      let instruction = elements[0];
      let value = parseInt(elements[1]);
      switch (instruction) {
        case 'acc':
          accumulator += value;
          i++;
          break;
        case 'jmp':
          i += value;
          break;
        default:
          i++;
      }
    }
    if (fixed) {
      break;
    }
    inputArray[j] = original;
  }
  return fixed;
}
if (isFix(inputArray, jmps, 'nop')) {
  console.log('step2:', accumulator);
}
if (isFix(inputArray, nops, 'jmp')) {
  console.log('step2:', accumulator);
}
console.log('end');
// function step2(inputArray: string[]) {

// }

// console.log(step2(inputArray));
