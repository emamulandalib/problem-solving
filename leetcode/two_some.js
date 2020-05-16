/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    for(let i in nums) {
        for (let j in nums) {
            if (i === j) continue

            if((nums[i] + nums[j]) === target) {
                return [Number(i), Number(j)]
            }
        }
    }

    return
};

console.log(twoSum([-18, 12, 3, 0], -6))