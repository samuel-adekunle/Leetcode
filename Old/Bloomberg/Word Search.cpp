#include <vector>
#include <string>
#include <iostream>
#include <unordered_map>

using namespace std;

bool helper(const vector<vector<char>>& board, const string& target, int i, int j, const string s, unordered_map<int, unordered_map<int, bool>> visited) {
  if (target == s) return true;
  if (i < 0 || i >= board.size() || j < 0 || j >= board[0].size()) {
    return false;
  }
  if (visited[i][j]) return false;
  size_t index = s.size();
  char curr = board[i][j];
  if (curr != target[index]) return false;
  visited[i][j] = true;
  const string _s = s + curr;
  return helper(board, target, i - 1, j, _s, visited)
    || helper(board, target, i + 1, j, _s, visited)
    || helper(board, target, i, j - 1, _s, visited)
    || helper(board, target, i, j + 1, _s, visited);
}

bool wordSearch(const vector<vector<char>>& board, const string& target) {
  unordered_map<int, unordered_map<int, bool>> visited;
  for (size_t i = 0; i < board.size(); i++) {
    for (size_t j = 0; j < board[0].size(); j++) {
      if (board[i][j] == target.front()) {
        if (helper(board, target, i, j, string(""), visited)) return true;
      }
    }
  }
  return false;
}

int main(int argc, char const* argv[])
{
  vector<vector<char>> board = {
    {'a', 'b', 'c', 'e'},
    {'s', 'd', 'c', 's'},
    {'a', 'd', 'e', 'e'}
  };
  string target = "abccedasa";
  cout << (wordSearch(board, target) ? "TRUE" : "FALSE") << endl;
  return 0;
}
