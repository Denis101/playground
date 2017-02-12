class QueueNode {
  constructor(value) {
    this.value = value;
    this.min = value;
    this.max = value;
    this.next = null;
  }
}

class Queue {
  constructor() {
    this.head = null;
    this.tail = null;
  }

  queue(value) {
    let node = new QueueNode(value);

    if (this.head == null) {
      this.head = node;

      if (this.tail == null) {
        this.tail = this.head;
      }

      return;
    }

    node.min = value < this.head.min ? value : this.head.min;
    node.max = value > this.head.max ? value : this.head.max;

    let tmp = this.head;
    tmp.next = node;
    this.head = node;

    if (this.tail == null) {
      this.tail = this.head;
    }
  }

  dequeue() {
    if (this.tail == null) {
      return null;
    }

    let tmp = this.tail;
    this.tail = tmp.next;
    return tmp.value;
  }

  min() {
    return this.head ? this.head.min : -1;
  }

  max() {
    return this.head ? this.head.max : -1;
  }
}

var queue = new Queue();
queue.queue(1);
queue.queue(3);
queue.queue(2);
console.log(queue.min());
console.log(queue.max());
console.log('dequeue!', queue.dequeue());
console.log(queue.min());
console.log(queue.max());
console.log('dequeue!', queue.dequeue());
console.log(queue.min());
console.log(queue.max());
console.log('dequeue!', queue.dequeue());
