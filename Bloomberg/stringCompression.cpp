/* Given an array of characters chars, compress it using the following algorithm:

Begin with an empty string s. For each group of consecutive repeating characters in chars:

If the group's length is 1, append the character to s.
Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

After you are done modifying the input array, return the new length of the array. */

#include <vector>
#include <string>

using namespace std;

// uses O(n) space
int compress(vector<char>& chars) {
  if (chars.size() < 2) {
    return chars.size();
  }

  char cur = chars[0];
  int cnt = 0;

  vector<char> compressed;
  compressed.reserve(chars.size());

  auto add_char = [&](const char c) {
    compressed.push_back(c);
    if (cnt > 1) {
      string s_cnt = to_string(cnt);
      for (const char s : s_cnt) {
        compressed.push_back(s);
      }
    }
  };

  for (const char c : chars) {
    if (c == cur) {
      cnt++;
      continue;
    }
    add_char(cur);
    cur = c;
    cnt = 1;
  }
  add_char(cur);

  chars.assign(compressed.begin(), compressed.end());

  return compressed.size();
}

// O(1) space => optimal solution because of BSR
int compress_linear(vector<char>& chars) {
  if (chars.size() < 2) {
    return chars.size();
  }

  char cur = chars[0];
  int cnt = 0;
  int length = 0;

  auto add_char = [&](const char c) {
    chars[length++] = c;
    if (cnt > 1) {
      string s_cnt = to_string(cnt);
      for (const char s : s_cnt) {
        chars[length++] = s;
      }
    }
  };

  for (const char c : chars) {
    if (c == cur) {
      cnt++;
      continue;
    }
    add_char(cur);
    cur = c;
    cnt = 1;
  }
  add_char(cur);

  return length;
}