#include <unordered_set>
#include <iostream>
#include <iterator>
#include <random>

class RandomizedSet
{
private:
  std::unordered_set<int> mem;
public:
  RandomizedSet();
  ~RandomizedSet();

  bool insert(int val);
  bool remove(int val);
  int getRandom() const;
};

RandomizedSet::RandomizedSet()
{
  srand(time(NULL));
}

RandomizedSet::~RandomizedSet()
{
}

bool RandomizedSet::insert(int val) {
  auto it = mem.find(val);
  if (it != mem.end()) return false;
  mem.insert(val);
  return true;
}

bool RandomizedSet::remove(int val) {
  auto it = mem.find(val);
  if (it == mem.end()) return false;
  mem.erase(val);
  return true;
}

int RandomizedSet::getRandom() const {
  int randomIndex = random() % mem.size();
  auto it = mem.begin();
  std::advance(it, randomIndex);
  return *it;
}

int main(int argc, char const* argv[])
{
  RandomizedSet* obj = new RandomizedSet;

  std::cout << obj->insert(1) << std::endl;
  std::cout << obj->remove(2) << std::endl;
  std::cout << obj->insert(2) << std::endl;
  std::cout << obj->getRandom() << std::endl;
  std::cout << obj->remove(1) << std::endl;
  std::cout << obj->insert(2) << std::endl;
  std::cout << obj->getRandom() << std::endl;

  delete obj;
  return 0;
}
