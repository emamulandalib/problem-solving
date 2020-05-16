

// Complete the reversePrint function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode next;
 * }
 *
 */
function reversePrint(head) {
    if (!head) {
        console.log(head)
        return
    }

    if (!head.next) {
        console.log(head.data)
        return
    }

    let data = []

    while(head) {
        data.unshift(head.data)
        head = head.next
    }
    
    data.forEach(v => console.log(v) + '\n')
}

