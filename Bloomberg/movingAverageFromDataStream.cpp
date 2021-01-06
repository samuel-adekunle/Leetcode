#include <deque>
#include <vector>
#include <iostream>

class MovingAverage {
private:
  std::deque<int> mem;
  size_t size;

  void addElement(const int i) {
    if (mem.size() == size) mem.pop_front();
    mem.push_back(i);
  }

  double average() const {
    double acc = 0;
    for (const int i : mem) acc += i;

    return acc / double(mem.size());
  }

public:
  MovingAverage(size_t size) : size(size) {}
  double next(const int);
};

double MovingAverage::next(const int i) {
  addElement(i);
  return average();
}

int main(int argc, char const* argv[])
{
  MovingAverage mavg(3);
  std::vector<int> arr = { 1,10,3,5 };

  for (const int i : arr) {
    std::cout << mavg.next(i) << std::endl;
  }

  return 0;
}