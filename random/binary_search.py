import typing


def binary_search(nums: typing.List[int], target: int) -> int:
    l, h = 0, len(nums) - 1

    while l < h:
        mid = (l + h) // 2

        if target == nums[mid]:
            return mid

        if target > nums[mid]:
            l = mid + 1

        if target < nums[mid]:
            h = mid - 1

    return -1


print(binary_search([2, 4, 6, 6], 6))
