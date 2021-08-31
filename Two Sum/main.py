from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(0, len(nums)):
            x = target - nums[i]
            try:
                attr_index = nums.index(x)
                if attr_index != i:
                    return i, attr_index
            except:
                pass

    def twoSumCached(self, nums: List[int], target: int) -> List[int]:
        cache = dict()
        for index, num in enumerate(nums):
            if (remainder := target - num) in cache:
                return index, cache[remainder]
            cache[num] = index


if __name__ == "__main__":
    solution = Solution()
    nums, target = [0, 4, 3, 0], 0

    print(solution.twoSum(nums, target))
    print(solution.twoSumCached(nums, target))