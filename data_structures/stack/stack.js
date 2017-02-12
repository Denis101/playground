class StackNode {
  constructor(value) {
    this.value = value;
    this.min = value;
    this.max = value;
    this.next = null;
  }
}

class Stack {
  constructor() {
    this.head = null;
  }

  push(value) {
    let node = new StackNode(value);

    if (this.head == null) {
      this.head = node;
      return;
    }

    node.min = value < this.head.min ? value : this.head.min;
    node.max = value > this.head.max ? value : this.head.max;

    let tmp = this.head;
    node.next = tmp;
    this.head = node;
  }

  pop() {
    if (this.head == null) {
      return null;
    }

    let tmp = this.head;
    this.head = tmp.next;
    return tmp.value;
  }

  min() {
    return this.head ? this.head.min : -1;
  }

  max() {
    return this.head ? this.head.max : -1;
  }
}

var stack = new Stack();
stack.push(3);
stack.push(2);
stack.push(1);
console.log(stack.min());
console.log(stack.max());
console.log(stack.pop());
console.log(stack.min());
console.log(stack.max());
