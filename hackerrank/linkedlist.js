const SinglyLinkedListNode = class {
    constructor(nodeData) {
        this.data = nodeData;
        this.next = null;
    }
};

const SinglyLinkedList = class {
    constructor() {
        this.head = null;
    }

};

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

    while (next) {
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

function reverse(head) {
    if (!head || !head.next) {
        return head
    }

    let newList = new SinglyLinkedListNode(head.data)


    while (head.next) {
        let newHead = new SinglyLinkedListNode(head.next.data)
        newHead.next = newList
        newList = newHead
        head = head.next
    }

    return newList
}

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

    while (head) {
        data.unshift(head.data)
        head = head.next
    }

    data.forEach(v => console.log(v) + '\n')
}

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

function main() {
    const llist = { "data": 1, "next": { "data": 2, next: null } }
    const llist2 = { "data": 1, "next": { "data": 3, next: null } }
    const llistResult = mergeLists(llist, llist2)
    console.log(JSON.stringify(llistResult))
}


main()
