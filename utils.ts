const fs = require('fs');
const os = require('os');

// inputString returns the input as a string
function inputString(filename:string) {
    return fs.readFileSync(filename,'utf8')
}

// inputStringArray returns the input split by new line
function inputStringArray(input:string) {
    return input.split(os.EOL)
}

// stringArrayToNumber converts the array of strings to numbers
function stringArrayToNumber(input:string[]) {
  let numbers = []
  input.forEach((s) => {
numbers.push(parseInt(s))
})
return numbers
}

module.exports = { inputString,inputStringArray,stringArrayToNumber}
