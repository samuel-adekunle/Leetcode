#include <vector>
#include <iostream>
#include <unordered_set>
#include <utility>

bool commonElements(std::vector<int>& arr1, std::vector<int>& arr2) {
  auto [_arr1, _arr2] = arr1.size() < arr2.size()
    ? std::pair<std::vector<int>&, std::vector<int>&>(arr1, arr2)
    : std::pair<std::vector<int>&, std::vector<int>&>(arr2, arr1);

  std::unordered_set<int> mem(_arr1.begin(), _arr1.end());
  for (const int i : _arr2) {
    if (mem.find(i) != mem.end()) return true;
  }
  return false;
}

int main(int argc, char const* argv[])
{
  std::vector<int> arr1 = { 1,2,3,4,5 };
  std::vector<int> arr2 = { 6,7,8 };

  std::cout << (commonElements(arr1, arr2) ? "TRUE" : "FALSE") << std::endl;

  return 0;
}
