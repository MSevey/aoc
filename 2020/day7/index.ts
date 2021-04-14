const utils = require('../../utils.ts');

// Grab the input
// let input = utils.inputString('input.txt');
let input = utils.inputString('test.txt');
let inputArray = utils.inputStringArray(input);

interface Bag {
  color: string;
  subBags: Map<Bag, number>;
}

function step1(inputArray: string[]) {
  // Parse input
  //
}
console.log(step1(inputArray));

// function step2(inputArray: string[]) {

// }

// console.log(step2(inputArray));
