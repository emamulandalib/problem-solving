import typing


def merge_sort(items: typing.List[int]) -> typing.List[int]:
    if len(items) < 2:
        return items

    f = merge_sort(items[:len(items) // 2])
    s = merge_sort(items[len(items) // 2:])
    return merge(f, s)


def merge(a: typing.List[int], b: typing.List[int]) -> typing.List[int]:
    result = []
    i, j = 0, 0

    while i < len(a) and j < len(b):
        if a[i] < b[j]:
            result.append(a[i])
            i += 1
        else:
            result.append(b[j])
            j += 1

    while i < len(a):
        result.append(a[i])
        i += 1

    while j < len(b):
        result.append(b[j])
        j += 1

    return result


print(merge_sort([3, 4, 1, 6]))
