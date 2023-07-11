## Linked List

A linked list is a data structure where data is stored in a sequence of nodes that are linked together. Unlike arrays, which store values in contiguous memory, linked lists store values in nodes that have references to the next node. The tutorial explains the difference between arrays and linked lists in terms of accessing and modifying values.

Adding or deleting an element at the beginning of a linked list is simple and takes constant time, regardless of the size of the list. This can be advantageous in certain applications. The tutorial also mentions that in arrays, when adding an element at the beginning, all the values need to shift, which can be more complicated than in a linked list.

The tutorial introduces two types of linked lists: singly linked lists and doubly linked lists. In a singly linked list, each node contains a reference to the next node. In a doubly linked list, each node contains references to both the next and previous nodes. The advantage of a doubly linked list is that adding an element at the end is easier than in a singly linked list.

The tutorial then demonstrates how to implement a linked list in Go. It defines a node struct that holds data and a pointer to the next node. It also defines a linked list struct that holds the head node and the length of the list. The tutorial provides a method called prepend, which adds a node to the front of the linked list.

Finally, a test is conducted by creating a linked list, adding nodes to it using the prepend method, and printing the list. The output shows the address of the head node and the length of the list.

In summary, the tutorial covers the basics of linked lists, explaining their advantages over arrays in certain scenarios and introducing both singly linked lists and doubly linked lists. The implementation of a linked list in Go is demonstrated through the creation of nodes and the prepend method.