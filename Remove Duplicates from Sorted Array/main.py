from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if (length := len(nums)) < 2:
            return length
        slowIndex = 0
        for fastIndex in range(0, length):
            slowElement, fastElement = nums[slowIndex], nums[fastIndex]
            if slowElement != fastElement:
                slowIndex += 1
                nums[slowIndex] = fastElement
        return slowIndex + 1


if __name__ == "__main__":
    solution = Solution()
    print(solution.removeDuplicates([1, 1, 2]))
