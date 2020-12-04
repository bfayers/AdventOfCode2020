const fs = require('fs');

//From https://stackoverflow.com/a/57969193
var lines = fs.readFileSync('./day4.txt', 'utf8').split('\n');

//Combine passports into one line each
passports = [];
for ( var i = 0; i<lines.length; i++ ) {
    newLine = "";
    let counter = 0;
    for ( var j = i; j<lines.length; j++ ) {
        //When the line goes blank
        if ( lines[j] != '' ) {
            counter++;
        } else {
            break;
        }
    }
    for ( var k = i; k<i+counter; k++ ) {
        newLine += " "+lines[k];
    }
    i+=counter;

    //Convert the combined line into a object.
    newPassport = {};
    newLine = newLine.substring(1,).split(" ");
    newLine = newLine.map(x => x.split(":"));
    //Makes the object
    newLine.forEach(x => {
        newPassport[x[0]] = x[1];
    });
    //Put into the passports array
    passports.push(newPassport);
    //Reset the newline var.
    newLine = "";
}

//Check passport validity.
validPassports = 0;

passports.forEach(passport => {
    //Check if all fields are present (except for cid)
    if ( typeof passport['byr'] !== 'undefined' && typeof passport['iyr'] !== 'undefined' &&  typeof passport['eyr'] !== 'undefined' && typeof passport['hgt'] !== 'undefined' && typeof passport['hcl'] !== 'undefined' && typeof passport['ecl'] !== 'undefined' && typeof passport['pid'] !== 'undefined' ) {
        validPassports++;
    }
});

console.log(validPassports);