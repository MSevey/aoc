const utils = require('../../utils.ts');

// Grab the input
let input = utils.inputString('input.txt');
let inputArray = utils.inputStringArray(input);

// Generate answer map
function step1(inputArray: string[]) {
  let totalAnswers = 0;
  let answerMap = new Map();
  inputArray.forEach((el) => {
    if (el == '') {
      totalAnswers += answerMap.size;
      answerMap.clear();
      return;
    }

    let letters = el.split('');
    letters.forEach((letter) => {
      if (answerMap.has(letter)) {
        return;
      }
      answerMap.set(letter, '');
    });
  });
  totalAnswers += answerMap.size;
  return totalAnswers;
}
console.log(step1(inputArray));

function step2(inputArray: string[]) {
  let totalAnswers = 0;
  let answerMap = new Map();
  let groupSize = 0;
  inputArray.forEach((el) => {
    if (el == '') {
      let size = 0;
      // TODO: Why doesn't map iteration work?
answerMap.forEach((value) => {
  if (value == groupSize) {
          size+=1;
        }
})
      totalAnswers += size;
      groupSize = 0;
      answerMap.clear();
      return;
    }
    groupSize+=1;
    let letters = el.split('');
    letters.forEach((letter) => {
      let old = answerMap.get(letter);
      if (old) {
        answerMap.set(letter, old+=1);
        return;
      }
      answerMap.set(letter, 1);
    });
  });
  let size = 0;
answerMap.forEach((value) => {
  if (value == groupSize) {
          size+=1;
        }
})

  totalAnswers += size;
  return totalAnswers;
}

function mapSize(answerMap: Map<string, number>, groupSize: number) {}
console.log(step2(inputArray));
