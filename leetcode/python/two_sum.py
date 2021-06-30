from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        pairs = {}

        for i, num in enumerate(nums):
            data = pairs.get(num)
            
            if data is not None:
                return [data, i]
            
            pairs[target - num] = i
        
        return []




nums = [2,7,11,15]
target = 9
result = Solution().twoSum(nums, target)
print(result)