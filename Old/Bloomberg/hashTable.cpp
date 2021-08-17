#include <functional>
#include <list>
#include <vector>
#include <string>
#include <iostream>
#include <algorithm>

template<class K, class V>
class HashTable {
private:
  size_t size;
  std::vector<std::list<std::pair<K, V>>> mem;

  size_t getHashIndex(const K& key) {
    size_t hashIndex = std::hash<K>{}(key);
    return hashIndex % size;
  }
public:
  HashTable(size_t size = 255) : size(size) {
    mem.resize(size);
  }

  void put(const K& key, const V& val) {
    std::list<std::pair<K, V>>& subList = mem[getHashIndex(key)];
    for (auto& kv : subList) {
      if (kv.first == key) { kv.second = val; return; }
    }
    mem[getHashIndex(key)].push_back(std::pair<K, V>(key, val));
  }

  void remove(const K& key) {
    std::list<std::pair<K, V>>& subList = mem[getHashIndex(key)];
    auto it = std::find_if(subList.begin(), subList.end(), [&key](const auto& kv) {
      return kv.first == key;
      });
    if (it != subList.end()) subList.erase(it);
  }

  V get(const K& key) {
    std::list<std::pair<K, V>>& subList = mem[getHashIndex(key)];
    for (const auto& kv : subList) {
      if (kv.first == key) return kv.second;
    }
    put(key, V());
    return V();
  }
};

int main(int argc, char const* argv[])
{
  HashTable<std::string, int> hashTable;
  hashTable.put("one", 1);
  hashTable.put("two", 2);
  std::cout << hashTable.get("one") << std::endl;
  std::cout << hashTable.get("three") << std::endl;
  hashTable.put("three", 3);
  std::cout << hashTable.get("three") << std::endl;
  hashTable.remove("one");
  std::cout << hashTable.get("one") << std::endl;
  return 0;
}
