#include <iostream>
#include <vector>
#include <unordered_set>
#include <algorithm>

class Solution {
public:
  int threeSumClosest(std::vector<int>& nums, const int target) {
    std::sort(nums.begin(), nums.end()); // O(n log n)
    int closest;
    for (size_t i = 0; i < nums.size() - 2; i++) { // O(n^2)
      int curr = nums[i] + twoSumClosest(nums, target - nums[i], i + 1);
      if (i == 0 || abs(target - curr) < abs(target - closest)) closest = curr;
      if (closest == target) break;
    }
    return closest;
  }

  // O(n)
  int twoSumClosest(const std::vector<int>& nums, const int target, int left) {
    int right = nums.size() - 1;
    int closest = nums[left] + nums[right];
    while (right > left) {
      int curr = nums[left] + nums[right];
      if (abs(target - curr) < abs(target - closest)) { closest = curr; }
      if (closest == target) break;
      if (closest > target) right--;
      else left++;
    }
    return closest;
  }
};

int main(int argc, char const* argv[])
{
  std::vector<int> nums = { 1,2,5,10,11,12 };
  const int target = 12;

  std::vector<int> nums1 = { -55,-24,-18,-11,-7,-3,4,5,6,9,11,23,33 };
  const int target1 = 0;

  Solution solution;

  std::cout << solution.threeSumClosest(nums, target) << std::endl;
  std::cout << solution.threeSumClosest(nums1, target1) << std::endl;
  return 0;
}
