const fs = require('fs');
const { parse } = require('path');

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

//Check passport validity. (Part 1)
validPassports = 0;

passports.forEach(passport => {
    //Check if all fields are present (except for cid)
    if ( typeof passport['byr'] !== 'undefined' && typeof passport['iyr'] !== 'undefined' &&  typeof passport['eyr'] !== 'undefined' && typeof passport['hgt'] !== 'undefined' && typeof passport['hcl'] !== 'undefined' && typeof passport['ecl'] !== 'undefined' && typeof passport['pid'] !== 'undefined' ) {
        validPassports++;
    }
});

console.log(validPassports);

//Check passport validity. (Part 2)
validPassports = 0;
passports.forEach(passport => {
    //Check if all fields are present (except for cid)
    if ( typeof passport['byr'] !== 'undefined' && typeof passport['iyr'] !== 'undefined' &&  typeof passport['eyr'] !== 'undefined' && typeof passport['hgt'] !== 'undefined' && typeof passport['hcl'] !== 'undefined' && typeof passport['ecl'] !== 'undefined' && typeof passport['pid'] !== 'undefined' ) {
        //check field validity
        let validity = 7;
        //Check Birth Year Validity more than 1920, less than 2002
        let byr = parseInt(passport['byr']);
        if ( ! (byr>=1920 && byr<=2002) ) {
            validity--;
        }
        //Check issue year more than 2010 less than 2020
        let iyr = parseInt(passport['iyr']);
        if ( ! (iyr>=2010 && iyr<=2020) ) {
            validity--;
        }
        //Check expiry year more than 2020 less than 2030
        let eyr = parseInt(passport['eyr']);
        if ( ! (eyr>=2020 && eyr<=2030) ) {
            validity--;
        }
        //Check height cm/in differs the amount
        let hgt = parseInt(passport['hgt']);       
        if ( passport['hgt'].includes('cm') ) {
            //It's in CM.
            if ( ! (hgt>=150 && hgt<=193) ) {
                validity--;
            }
        } else {
            //It's in in.
            if ( ! (hgt>=59 && hgt<=76) ) {
                validity--;
            }
        }
        //Validate hair color
        //Regex from https://mkyong.com/regular-expressions/how-to-validate-hex-color-code-with-regular-expression/
        let hairRe = new RegExp('^#([a-fA-F0-9]{6})$');
        if ( hairRe.exec(passport['hcl']) == null ) {
            validity--;
        }
        //Validate eye colour
        eyeColours = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];
        if ( ! eyeColours.includes(passport['ecl']) ) {
            validity--;
        }

        //Validate passport ID
        let pidRe = new RegExp('^\\d{9}$');
        if ( pidRe.exec(passport['pid']) == null ) {
            validity--;
        }

        if ( validity == 7 ) {
            validPassports++;
        }
    }
});

console.log(validPassports);
