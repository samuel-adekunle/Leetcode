from typing import List


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int],
              n: int) -> None:
        if n == 0:
            return
        if m == 0:
            for i in range(n):
                nums1[i] = nums2[i]
            return

        pm, pn = m - 1, n - 1
        for i in range(m + n - 1, -1, -1):
            print(nums1)
            xm, xn = nums1[pm], nums2[pn]
            if xm > xn:
                nums1[i] = xm
                pm -= 1
                if pm < 0: break
            else:
                nums1[i] = xn
                pn -= 1
                if pn < 0: return

        i -= 1
        while i >= 0 and pn >= 0:
            nums1[i] = nums2[pn]
            i -= 1
            pn -= 1


if __name__ == "__main__":
    solution = Solution()
    nums1, m = [2, 2, 7, 0, 0, 0], 3
    nums2, n = [5, 8, 10], 3
    solution.merge(nums1, m, nums2, n)
    print(nums1)