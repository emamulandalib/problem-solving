

// Complete the insertNodeAtTail function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode next;
 * }
 *
 */
function insertNodeAtTail(head, data) {
    let node = new SinglyLinkedListNode(data)
    
    if(!head) {
        return node
    } 

    let next = head

    while(next.next) {
        next = next.next
    }
    
    next.next = node
    return head
}

