

// Complete the CompareLists function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode next;
 * }
 *
 */
function CompareLists(llist1, llist2) {
    if (!llist1 && !llist2) {
        return 0
    }

    while (llist1) {
        let v2 = null
        let v1 = llist1.data

        while (llist2) {
            v2 = llist2.data
            llist2 = llist2.next
            break
        }

        if (v1 !== v2) return 0
        llist1 = llist1.next
    }

    return 1
}

