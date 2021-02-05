#include <string>
#include <vector>
#include <utility>
#include <cassert>
#include <iostream>

enum RegexType {
  Single,
  ZeroOrMore,
};

typedef std::pair<char, RegexType> TokenType;
constexpr char zeroOrMore = '*';
constexpr char anyChar = '.';

class SimpleRegexEngine {
private:
  const std::string ref;
  std::vector<TokenType> tokens;

  void tokenise() {
    auto it = ref.rbegin();
    TokenType token;
    while (it != ref.rend()) {
      switch (*it)
      {
      case zeroOrMore:
        std::advance(it, 1);
        token.second = ZeroOrMore;
        token.first = *it;
        break;
      default:
        token.first = *it;
        token.second = Single;
        break;
      }
      tokens.push_back(token);
      std::advance(it, 1);
    }
  }

  bool matchToken(const TokenType& token, const char c) const {
    return token.first == anyChar || c == token.first;
  }

public:
  SimpleRegexEngine(const std::string& ref) : ref(ref) {
    tokenise();
  }

  bool matchString(const std::string s) const {
    auto it = s.crbegin();
    for (const TokenType& token : tokens) {
      if (it == s.crend()) return false;
      switch (token.second)
      {
      case Single:
        if (!matchToken(token, *it)) return false;
        std::advance(it, 1);
        break;
      case ZeroOrMore:
        //! need to be able to recursively check all cases i.e. if zero or one or two
        //! currently just matching until not able to, so max possible matches
        while (it != s.crend() && matchToken(token, *it)) std::advance(it, 1);
        break;
      }
    }
    return it == s.crend();
  }
};

int main(int argc, char const* argv[])
{
  SimpleRegexEngine re = std::string("c*a*b");
  std::cout << re.matchString("aab") << std::endl;
  std::cout << re.matchString("abcl") << std::endl;
  std::cout << re.matchString("") << std::endl;
  std::cout << re.matchString("a") << std::endl;
  std::cout << re.matchString("acl") << std::endl;
  std::cout << re.matchString("ac") << std::endl;
  std::cout << re.matchString("abbbbbbbcl") << std::endl;
  return 0;
}
