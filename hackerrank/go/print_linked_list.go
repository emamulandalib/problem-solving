

// Complete the printLinkedList function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
 func printLinkedList(head *SinglyLinkedListNode) {
    fmt.Println(head.data)

    if head.next != nil {
        printLinkedList(head.next)
    }
}

