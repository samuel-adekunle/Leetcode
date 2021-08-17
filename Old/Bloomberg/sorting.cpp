#include <vector>
#include <functional>
#include <iostream>
#include <cassert>

template<class T>
class VectorSort {
private:
  void _swap(std::vector<T>& arr, size_t i, size_t j) {
    T tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;
  }

  std::vector<T> mergeSorted(const std::vector<T>& arr1, const std::vector<T>& arr2) {
    std::vector<T> arr;
    arr.reserve(arr1.size() + arr2.size());
    size_t i = 0, j = 0;

    while (i < arr1.size() && j < arr2.size()) {
      if (arr1[i] < arr2[j]) arr.push_back(arr1[i++]);
      else arr.push_back(arr2[j++]);
    }

    if (i < arr1.size()) arr.insert(arr.begin() + i + j, arr1.begin() + i, arr1.end());
    else arr.insert(arr.begin() + i + j, arr2.begin() + j, arr2.end());

    return arr;
  }

  std::vector<T> mergeHelper(const std::vector<T>& arr, size_t i, size_t j) {
    if (j - i == 1) return { arr[i] };
    size_t mid = i + (j - i) / 2;
    std::vector<T> left = mergeHelper(arr, i, mid);
    std::vector<T> right = mergeHelper(arr, mid, j);
    return mergeSorted(left, right);
  }

  // { 6,5,3,1,8,7,2,7,9,-1,0,3 };

  void quickHelper(std::vector<T>& arr, size_t i, size_t j) {
    const size_t _i = i, _j = j;
    if (i >= j || j > arr.size()) return;
    T pivot = arr[j - 1];
    while (i < j) {
      if (arr[i] <= pivot) {
        i++;
        continue;
      }
      _swap(arr, j - 2, j - 1);
      _swap(arr, i, j - 1);
      j--;
    }

    // for (const int i : arr) { std::cerr << i << " "; }
    // std::cerr << std::endl;
    // std::cerr << i << " " << j << " " << arr[j] << std::endl;
    // exit(EXIT_FAILURE);

    quickHelper(arr, _i, i + 1);
    quickHelper(arr, i + 2, _j);
  }

public:
  // bubble up the largest value
  void bubble(std::vector<T>& arr) {
    bool stable;
    for (size_t i = 0; i < arr.size(); i++) {
      stable = true;
      for (size_t j = 0; j < arr.size() - 1; j++) {
        if (arr[j] > arr[j + 1]) {
          _swap(arr, j, j + 1);
          stable = false;
        }
      }
      if (stable) break;
    }
  }

  // maintain ordered sub array
  void insertion(std::vector<T>& arr) {
    for (size_t i = 0; i < arr.size(); i++) {
      T curr = arr[i];
      for (size_t j = i - 1; j >= 0 && arr[j] > curr; j--) {
        _swap(arr, j, j + 1);
      }
    }
  }

  // bubble down smallest value
  void selection(std::vector<T>& arr) {
    T localMinVal;
    size_t localMinIndex;
    for (size_t i = 0; i < arr.size(); i++) {
      localMinVal = arr[i];
      localMinIndex = i;
      for (size_t j = i + 1; j < arr.size(); j++) {
        if (arr[j] < localMinVal) {
          localMinVal = arr[j];
          localMinIndex = j;
        }
      }
      if (localMinIndex != i) _swap(arr, i, localMinIndex);
    }
  }

  // divide and conquer approach
  void merge(std::vector<T>& arr) {
    arr = mergeHelper(arr, 0, arr.size());
  }

  // pick a pivot (the last element) and rotate around element
  void quick(std::vector<T>& arr) {
    quickHelper(arr, 0, arr.size());
  }

};

int main(int argc, char const* argv[])
{
  std::vector<int> arr = { 6,5,3,1,8,7,2,7,9,-1,0,3 };
  VectorSort<int> sorting;
  // sorting.bubble(arr);
  // sorting.selection(arr);
  // sorting.insertion(arr);
  // sorting.merge(arr);
  sorting.quick(arr);

  for (const int i : arr) std::cout << i << " ";
  std::cout << std::endl;
  return 0;
}
