

// Complete the reverse function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode next;
 * }
 *
 */
function reverse(head) {
    if (!head || !head.next) {
        return head
    }

    let newList = new SinglyLinkedListNode(head.data)


    while(head.next) {
        let newHead = new SinglyLinkedListNode(head.next.data)
        newHead.next = newList
        newList = newHead
        head = head.next
    }
    
    return newList
}

