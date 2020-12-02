const { count, countReset } = require('console');
const fs = require('fs');

//From https://stackoverflow.com/a/57969193
var lines = fs.readFileSync('./day2.txt', 'utf8').split('\n');

//Make a counter to store number of valid PWs
var counter = 0;

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
        counter++;
    }
}


lines.forEach(part1);

console.log(counter);