const assert = require("assert");

const chop = (needle: number, haystack: number[]): number => {
  let currentHaystack = haystack;
  let offset = 0;

  while (currentHaystack.length !== 0) {
    const index = getIndex(currentHaystack);
    const element = currentHaystack[index];
    if (element === needle) return index + offset;
    if (element < needle) {
      currentHaystack = currentHaystack.slice(index + 1);
      offset += index + 1;
    } else {
      currentHaystack = currentHaystack.slice(0, index);
    }
  }
  return -1;
};

const getIndex = (haystack: number[]) => {
  return Math.floor(haystack.length / 2);
};

const test_chop = () => {
  assert.equal(-1, chop(3, []));
  assert.equal(-1, chop(3, [1]));
  assert.equal(0, chop(1, [1]));
  assert.equal(0, chop(1, [1, 3, 5]));
  assert.equal(1, chop(3, [1, 3, 5]));
  assert.equal(2, chop(5, [1, 3, 5]));
  assert.equal(-1, chop(0, [1, 3, 5]));
  assert.equal(-1, chop(2, [1, 3, 5]));
  assert.equal(-1, chop(4, [1, 3, 5]));
  assert.equal(-1, chop(6, [1, 3, 5]));
  assert.equal(0, chop(1, [1, 3, 5, 7]));
  assert.equal(1, chop(3, [1, 3, 5, 7]));
  assert.equal(2, chop(5, [1, 3, 5, 7]));
  assert.equal(3, chop(7, [1, 3, 5, 7]));
  assert.equal(-1, chop(0, [1, 3, 5, 7]));
  assert.equal(-1, chop(2, [1, 3, 5, 7]));
  assert.equal(-1, chop(4, [1, 3, 5, 7]));
  assert.equal(-1, chop(6, [1, 3, 5, 7]));
  assert.equal(-1, chop(8, [1, 3, 5, 7]));
};

test_chop();
console.log("Success!");
