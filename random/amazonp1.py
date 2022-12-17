# l = [2 ,3, 4, 3, 3]
# [0,1,2,1,1]
# [0,0,1,0,0]
# [0,0,0,0,0]


# l = [3,2,4,3,3,3]
# l = [0, 1, 1, 0, 0, 0]
# l = [0, 0, 0, 0, 0, 0]

# l = [3,2,4,3,3,3]
# l = [1, 0, 2, 1, 1]
# l =

# 1 < n < 10^4

# min = min(l) = 2 O(n)
# l = [2,3,3,3,4]
# l = [0,1,1,1,2]
# d = {
# 0: 2
# 1: 3,
# 2: 4,
# 3: 3,
# 4, 3
# }


res = 0
def main():
    input = [2 ,3, 4, 3, 3]
    input.sort() # n log n = [2,3,3,4]
    sol(input)
    print(0)

def sol(l):
    if len(l) == 0:
        return

    res = res + 1
    min_v  = l[0]

    if min_v == 0:
        return sol(l[1:])

    for i in range(len(l)):
        l[i] = l[i] - min_v

    return sol(l)

main()