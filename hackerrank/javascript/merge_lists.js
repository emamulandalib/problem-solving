

// Complete the mergeLists function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode next;
 * }
 *
 */
function mergeLists(head1, head2) {
    if (!head1 && !head2) return null

    let data = []

    while(head1) {
        data.push(head1.data)
        head1 = head1.next
    }

    while(head2) {
        data.push(head2.data)
        head2 = head2.next
    }

    let newHead = null

    data.sort((a,b) => a-b).forEach(v => {
        let node = new SinglyLinkedListNode(v)

        if (!newHead) {
            newHead = node 
            return
        }

        let next = newHead

        while(next.next) {
            next = next.next
        }

        next.next = node
    })

    return newHead

}

