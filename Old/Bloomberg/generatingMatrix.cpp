#include <vector>
#include <iostream>

using namespace std;

vector<vector<int>> generateMatrix(int n) {
  vector<vector<int>> spiral = vector<vector<int>>(n, vector<int>(n, 0)); // reserve the needed space;
  int x1 = 0, y1 = 0;
  int x2 = n - 1, y2 = n - 1;
  int i = 1;
  while (i <= n * n) {
    for (int a = x1, b = y1; b <= y2; b++) spiral[a][b] = i++;
    for (int a = x1 + 1, b = y2; a <= x2; a++) spiral[a][b] = i++;
    for (int a = x2, b = y2 - 1; b >= y1; b--) spiral[a][b] = i++;
    for (int a = x2 - 1, b = y1; a > x1; a--) spiral[a][b] = i++;
    x1++; y1++;
    x2--; y2--;
  }
  return spiral;
}

int main(int argc, char const* argv[])
{
  auto mat = generateMatrix(0);
  for (auto sub : mat) {
    for (auto i : sub) cout << i << " ";
    cout << endl;
  }
  return 0;
}
