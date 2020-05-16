

// Complete the deleteNode function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode next;
 * }
 *
 */
function deleteNode(head, position) {
    let count = 0

    if (!head) {
        return head
    }

    if (position === count) {
        return head ? head.next : null
    }

    let previous = head
    let next = head.next

    while(next) {
        count++

        if (position === count) {
            previous.next = next.next
            return head
        }

        previous = next
        next = next.next
    }
    
    return head
}

