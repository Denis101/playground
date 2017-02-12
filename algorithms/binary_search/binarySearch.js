const search = (target, values) => {
  console.log('SEARCHING FOR ' + target);
  
  let min = 0;
  let max = values.length;

  let i = 0;
  while (min <= max) {
    let current = Math.floor((min + max) / 2);
    i++;

    if (values[current] < target) {
      min = current + 1;
    } else if (values[current] > target) {
      max = current - 1;
    } else {
      console.log('Found ' + values[current] + ' in ' + i + ' attempts. O(' + (i / values.length) + ')');
      console.log(i + ' ITERATIONS');
      console.log(values.length + ' LENGTH');

      return current;
    }
  }

  return -1;
};

console.log('RESULT: ', search(0, [0]));
console.log('RESULT: ', search(1, [0, 1, 2, 3, 4, 5]));
console.log('RESULT: ', search(4, [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]));
console.log('RESULT: ', search(25, [11, 25, 48, 52, 78, 99, 101, 173, 256, 567, 1024, 1564, 3255, 4387, 11011, 13245, 23456, 101239]));
console.log('RESULT: ', search(99, [11, 25, 48, 52, 78, 99, 101, 173, 256, 567, 1024, 1564, 3255, 4387, 11011, 13245, 23456, 101239]));
console.log('RESULT: ', search(3255, [11, 25, 48, 52, 78, 99, 101, 173, 256, 567, 1024, 1564, 3255, 4387, 11011, 13245, 23456, 101239]));
console.log('RESULT: ', search(11011, [11, 25, 48, 52, 78, 99, 101, 173, 256, 567, 1024, 1564, 3255, 4387, 11011, 13245, 23456, 101239]));
console.log('RESULT: ', search(101239, [11, 25, 48, 52, 78, 99, 101, 173, 256, 567, 1024, 1564, 3255, 4387, 11011, 13245, 23456, 101239]));
