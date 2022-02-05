import collections


def bfs(graph, start):
    queue = collections.deque([start])
    visited = {start}

    while queue:
        v = queue.popleft()
        print(v)

        for n in graph[v] - visited:
            visited.add(n)
            queue.append(n)


g = {0: {1, 2}, 1: {2}, 2: {3}, 3: {1, 2}}
bfs(g, 0)
