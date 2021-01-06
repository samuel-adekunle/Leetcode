#include <vector>
#include <string>
#include <unordered_set>
#include <iostream>
#include <sstream>

using namespace std;

string filter(const string& s, char val) {
  stringstream filtered;

  for (const char c : s) {
    if (c != val) filtered << c;
  }

  return filtered.str();
}

string applyRules(const string& s) {
  size_t index = s.find('@');
  string domainName = s.substr(index);
  string localName = s.substr(0, index);

  // plus rule
  index = localName.find('+'); // find first occurrence of '+'
  localName = localName.substr(0, index); // and take everything before it

  // dot rule
  localName = filter(localName, '.');
  return localName + domainName;
}

int numUniqueEmails(const vector<string>& emails) {
  unordered_set<string> uniqueEmails;

  for (const string& email : emails) {
    uniqueEmails.insert(applyRules(email));
  }

  return uniqueEmails.size();
}

int main(int argc, char const* argv[])
{
  const vector<string> emails = {
    "test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"
  };

  cout << numUniqueEmails(emails) << endl;
  return 0;
}

