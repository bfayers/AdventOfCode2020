const { count, countReset } = require('console');
const fs = require('fs');

//From https://stackoverflow.com/a/57969193
var lines = fs.readFileSync('./day2.txt', 'utf8').split('\n');

//Make a counter to store number of valid PWs
var part1Counter = 0;

//Use forEach to go over every element.
function part1(element) {
    let policy = element.split(":")[0];
    let password = element.split(": ")[1];
    let policyMin = parseInt(policy.split("-")[0]);
    let policyMax = parseInt(policy.split("-")[1].split(" ")[0]);
    let policyLetter = policy.split("-")[1].split(" ")[1];

    //Count letter occurrences
    let countOfLetter = password.split(policyLetter).length-1;
    if ( policyMin <= countOfLetter && policyMax >= countOfLetter ) {
        //It's valid
        part1Counter++;
    }
}

//Make a counter to store number of valid PWs
var part2Counter = 0;

function part2(element) {
    let policy = element.split(":")[0];
    let password = element.split(": ")[1];
    let policyPos1 = parseInt(policy.split("-")[0])-1;
    let policyPos2 = parseInt(policy.split("-")[1].split(" ")[0])-1;
    let policyLetter = policy.split("-")[1].split(" ")[1];
    

    //Split password to letters
    password = password.split("");
    if ( password[policyPos1] == policyLetter ^ password[policyPos2] == policyLetter ) {
        part2Counter++;
    }
}


lines.forEach(part1);
lines.forEach(part2);

console.log(part1Counter);
console.log(part2Counter);