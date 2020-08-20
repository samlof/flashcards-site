const numbers = [
  "nolla",
  "yksi",
  "kaksi",
  "kolme",
  "neljä",
  "viisi",
  "kuusi",
  "seitsemän",
  "kahdeksan",
  "yhdeksän",
  "kymmenen",
];

export const numberToString = (number: number): string => {
  if (!number) return numbers[0];
  let ret = "";
  if (number > 99) {
    const hundreds = Math.floor(number / 100);
    if (hundreds === 1) ret += "sata";
    else ret += numbers[hundreds] + "sataa";
    number -= 100 * hundreds;
  }
  if (number > 19) {
    const tens = Math.floor(number / 10);
    const ones = number % 10;
    ret += numbers[tens] + "kymmentä";
    if (ones) ret += numbers[ones];
    return ret;
  }
  if (number > 10) {
    return ret + numbers[number - 10] + "toista";
  }
  return ret + numbers[number];
};

const checkSingleNumber = (str: string): [number, string] | undefined => {
  for (let i = 0; i < numbers.length; i++) {
    const num = numbers[i];

    if (str.startsWith(num)) {
      // save
      return [i, num];
    }
  }
};

const stringsToCheck: [string, number, string?][] = [
  ["sataa", 100],
  ["sata", 100],
  ["toista", 10, "+"],
  ["kymmentä", 10],
  ["kymmenen", 10],
];

export const stringToNumber = (numStr: string): number | undefined => {
  numStr = numStr.toLowerCase();
  const originalString = numStr;

  let result: number | undefined;
  while (numStr) {
    const firstNumArray = checkSingleNumber(numStr);
    let found = false;
    let foundNumber = 1;
    if (firstNumArray) {
      found = true;
      foundNumber = firstNumArray[0];
      numStr = numStr.substr(firstNumArray[1].length);
    }

    for (const checkPair of stringsToCheck) {
      if (numStr.startsWith(checkPair[0])) {
        found = true;
        if (checkPair[2] === "+") foundNumber += checkPair[1];
        else foundNumber *= checkPair[1];
        numStr = numStr.substr(checkPair[0].length);
        break;
      }
    }

    if (!result) result = 0;
    if (!found) {
      const asString = numberToString(result);
      if (originalString !== asString) {
        return undefined;
      }
      return result;
    }
    result += foundNumber;
  }

  // validate it's a good number
  if (result) {
    const asString = numberToString(result);
    if (originalString !== asString) {
      return undefined;
    }
  }
  return result;
};

const test = () => {
  for (let i = 0; i < 1000; i++) {
    const asString = numberToString(i);
    const backToNumber = stringToNumber(asString);

    if (i !== backToNumber) {
      throw new Error(`Expected ${i} but got ${backToNumber}. (${asString})`);
    }
  }
};

test();
