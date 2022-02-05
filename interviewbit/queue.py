from collections import deque


def queue():
    q = deque()
    q.append(1)
    q.append(2)
    q.append(3)
    print(q)

    q.pop()
    print(q)


queue()
