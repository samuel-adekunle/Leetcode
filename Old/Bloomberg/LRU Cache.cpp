#include <algorithm>
#include <unordered_map>
#include <list>
#include <iostream>

class LRUCache {
private:
  size_t capacity;
  std::unordered_map<int, int> mem;
  std::list<int> cache;

  // O(n)
  bool moveToFront(const int key) {
    auto it = std::find(cache.begin(), cache.end(), key); // O(n)
    if (it == cache.end()) {
      cache.push_front(key);
      return false;
    }
    if (it != cache.begin())
      cache.splice(cache.begin(), cache, it, std::next(it));
    return true;
  }


public:
  LRUCache(size_t capacity) : capacity(capacity) {
    mem.reserve(capacity);
  }

  // O(n)
  int get(int key) {
    auto it = mem.find(key);
    if (it == mem.end()) return -1;
    moveToFront(key); // O(n)
    return it->second;
  }

  // O(n)
  void put(int key, int value) {
    if (!moveToFront(key) && (mem.size() == capacity)) {
      int lruKey = cache.back();
      cache.remove(lruKey); // O(n)
      mem.erase(lruKey);
    }
    mem[key] = value;

  }
};

int main(int argc, char const* argv[])
{
  LRUCache lRUCache(2);
  lRUCache.put(1, 1); // cache is {1=1}
  lRUCache.put(2, 2); // cache is {1=1, 2=2}
  std::cout << lRUCache.get(1) << std::endl;;    // return 1
  lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
  std::cout << lRUCache.get(2) << std::endl;    // returns -1 (not found)
  lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
  std::cout << lRUCache.get(1) << std::endl;    // return -1 (not found)
  std::cout << lRUCache.get(3) << std::endl;    // return 3
  std::cout << lRUCache.get(4) << std::endl;    // return 4

  return 0;
}
