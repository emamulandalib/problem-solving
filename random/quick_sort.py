import typing


def quick_sort(items: typing.List[int]) -> typing.List[int]:
    if len(items) < 2:
        return items

    return quick_sort_run(items, 0, len(items) - 1)


def quick_sort_run(arr: typing.List[int], low: int, high: int) -> typing.List[int]:
    if low < high:
        arr, p = partition(arr, low, high)
        arr = quick_sort_run(arr, low, p - 1)
        arr = quick_sort_run(arr, p + 1, high)
    return arr


def partition(arr: typing.List[int], low: int, high: int) -> (typing.List[int], int):
    pivot = arr[high]
    i = low

    for j in range(low, high):
        if arr[j] < pivot:
            arr[i], arr[j] = arr[j], arr[i]
            i += 1

    arr[i], arr[high] = arr[high], arr[i]
    return arr, i


print(quick_sort([9, 23, 4, 4234, 4, 4, 4]))
