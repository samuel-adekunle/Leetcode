#include <vector>
#include <unordered_set>
#include <iostream>

using namespace std;

// O(m*n) space usage for visited array
vector<int> spiralOrder(vector<vector<int>>& matrix) {
  vector<int> spiral;
  if (matrix.size() == 0) return spiral;
  int m = matrix.size();
  int n = matrix[0].size();
  spiral.reserve(m * n); // DANGER - empty matrix
  bool visited[m][n];
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) {
      visited[i][j] = false;
    }
  }
  int i = 0;
  int j = -1;
  auto push_back = [&](int i, int j) {
    int cur = matrix[i][j];
    spiral.push_back(cur);
  };
  while (true) {
    // move right
    while (j < n - 1 && !visited[i][j + 1]) { push_back(i, ++j); visited[i][j] = true; }
    // move down
    while (i < m - 1 && !visited[i + 1][j]) { push_back(++i, j); visited[i][j] = true; }
    // move left
    while (j > 0 && !visited[i][j - 1]) { push_back(i, --j); visited[i][j] = true; }
    // move up
    while (i > 0 && !visited[i - 1][j]) { push_back(--i, j); visited[i][j] = true; }
    if (spiral.size() == spiral.capacity()) break;
  }
  return spiral;
}

/*
  using layers

 */
vector<int> spiralConstantMemory(vector<vector<int>>& matrix) {
  vector<int> spiral;
  if (matrix.size() == 0 || matrix[0].size() == 0) return spiral; // Danger: Short Circuit expression
  int m = matrix.size();
  int n = matrix[0].size();
  spiral.reserve(m * n); // DANGER - empty matrix
  int x1 = 0, y1 = 0;
  int x2 = m - 1, y2 = n - 1;

  while (x1 <= x2 && y1 <= y2) {
    // move right
    for (int i = x1, j = y1; j <= y2; j++) spiral.push_back(matrix[i][j]);
    // move down
    for (int i = x1 + 1, j = y2; i <= x2; i++) spiral.push_back(matrix[i][j]);
    if (y1 < y2 && x1 < x2) {
      // move left
      for (int i = x2, j = y2 - 1; j >= y1; j--) spiral.push_back(matrix[i][j]);
      // move up
      for (int i = x2 - 1, j = y1; i > x1; i--) spiral.push_back(matrix[i][j]);
    }

    x1++; y1++;
    x2--; y2--;
  }
  return spiral;
}

int main() {
  vector<vector<int>> mat = { { 1,2,3,4 }, { 5,6,7,8 }, { 9,10,11,12 } };
  vector<vector<int>> mat1 = { {1} };
  vector<int> res = spiralConstantMemory(mat);

  for (auto i : res) {
    cout << i << " ";
  }
  cout << endl;
}